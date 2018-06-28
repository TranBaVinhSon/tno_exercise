package main

import (
	"google.golang.org/grpc"
	"log"
	"github.com/tnakade/tno_exercise/app/proto/services"
	"context"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := services.NewWalletClient(conn)

	ctx := context.Background()
	req := services.GetBalanceRequest{}
	req.UserId = 1
	res, err := client.GetBalance(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Balance is %s", res.Balance)
}
