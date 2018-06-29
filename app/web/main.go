package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"github.com/tnakade/tno_exercise/app/proto/services"
	"log"
	"google.golang.org/grpc"
	"context"
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

	// ルーティング
	e.GET("/hello", HelloPage)
	e.GET("/shops/:id", ShopPage)
	e.GET("/shops/:id/balance", ShopBalancePage)

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
	templates["hello"] = template.Must(template.ParseFiles(baseTemplate, "templates/hello.html"))
	templates["shop"] = template.Must(template.ParseFiles(baseTemplate, "templates/shop.html"))
	templates["shop_balance"] = template.Must(template.ParseFiles(baseTemplate, "templates/shop_balance.html"))
}

func HelloPage(c echo.Context) error {
	greetingto := c.QueryParam("greetingto")
	return c.Render(http.StatusOK, "hello", greetingto)
}

func ShopPage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.Render(http.StatusOK, "shop", id)
}

func ShopBalancePage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("", id)

	req := services.GetBalanceRequest{UserId: uint64(id)}
	res, err := client.GetBalance(context.Background(), &req)

	if err != nil {
		return c.Render(http.StatusOK, "shop_balance", "0")
	}
	return c.Render(http.StatusOK, "shop_balance", res.Balance)
}
