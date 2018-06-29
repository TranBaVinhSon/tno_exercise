package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
	"github.com/tnakade/tno_exercise/app/proto/services"
	"log"
	"google.golang.org/grpc"
	"context"
	"strconv"
	"os"
	"image/png"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode"
	"fmt"
)

var templates map[string]*template.Template
var client services.WalletClient
// Template はHTMLテンプレートを利用するためのRenderer Interfaceです。
type Template struct {
}

// Render はHTMLテンプレートにデータを埋め込んだ結果をWriterに書き込みます。
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return templates[name].ExecuteTemplate(w, "layout.html", data)
}

func main() {
	conn, err := grpc.Dial("35.187.215.246:1080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client = services.NewWalletClient(conn)

	t := &Template{}

	// Echoのインスタンス作る
	e := echo.New()
	e.Renderer = t

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public", "public")

	// ルーティング
	e.GET("/shops/:id", ShopPage)
	e.GET("/shops/:id/balance", ShopBalancePage)
	e.GET("/shops/:id/transactions", ShopTransactionsPage)

	// サーバー起動
	e.Start(":2080") //ポート番号指定してね
}

// 初期化を行います。
func init() {
	loadTemplates()
}

// 各HTMLテンプレートに共通レイアウトを適用した結果を保存します（初期化時に実行）。
func loadTemplates() {
	var baseTemplate = "templates/layout.html"
	templates = make(map[string]*template.Template)
	templates["shop"] = template.Must(template.ParseFiles(baseTemplate, "templates/shop.html"))
	templates["shop_balance"] = template.Must(template.ParseFiles(baseTemplate, "templates/shop_balance.html"))
	templates["shop_transactions"] = template.Must(template.ParseFiles(baseTemplate, "templates/shop_transactions.html"))
}

func getId(c echo.Context) uint64 {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0
	}
	return uint64(id)
}

func createShopQRCodeFilePath(id uint64) string {
	return fmt.Sprintf("public/img/shop_%d.png", id)
}

func createShopQRCode(id uint64) {
	qrCode, _ := qr.Encode(strconv.FormatUint(id, 10), qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	fileName := createShopQRCodeFilePath(id)
	os.Remove(fileName)
	file, _ := os.Create(fileName)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}

type ShopPageInfo struct {
	Id uint64
	Path string
}

func ShopPage(c echo.Context) error {
	id := getId(c)
	createShopQRCode(id)
	info := ShopPageInfo{Id: id, Path: "/" + createShopQRCodeFilePath(id)}
	return c.Render(http.StatusOK, "shop", info)
}

func ShopBalancePage(c echo.Context) error {
	id := getId(c)

	req := services.GetBalanceRequest{UserId: uint64(id)}
	res, err := client.GetBalance(context.Background(), &req)

	if err != nil {
		return c.Render(http.StatusOK, "shop_balance", "0")
	}
	return c.Render(http.StatusOK, "shop_balance", res.Balance)
}

func ShopTransactionsPage(c echo.Context) error {
	id := getId(c)

	req := services.GetTransactionsRequest{UserId: id}
	res, err := client.GetTransactions(context.Background(), &req)

	if err != nil {
		res := services.GetTransactionsResponse{}
		res.Transactions = []*services.Transaction{}
		return c.Render(http.StatusOK, "shop_transactions", res)
	}
	return c.Render(http.StatusOK, "shop_transactions", res)
}
