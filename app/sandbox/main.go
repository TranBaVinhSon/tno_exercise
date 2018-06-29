package main

import (
	"context"
	"log"

	"github.com/tnakade/tno_exercise/app/proto/services"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := services.NewWalletClient(conn)

	callGetBalance(client)
	callGetTransactions(client)
	// callSendCoin(client)
}

func callGetBalance(client services.WalletClient) {
	ctx := context.Background()
	req := services.GetBalanceRequest{}
	req.UserId = 1
	res, err := client.GetBalance(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Balance is %s", res.Balance)
}

func callGetTransactions(client services.WalletClient) {
	ctx := context.Background()
	req := services.GetTransactionsRequest{}
	req.UserId = 1
	res, err := client.GetTransactions(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	for _, transaction := range res.Transactions {
		log.Printf(
			"Transaction: id[%s] category[%s] abandoned[%s] account[%s] address[%s] amount[%s]",
			transaction.Id,
			transaction.Category,
			transaction.Abandoned,
			transaction.ReceivedAccount,
			transaction.ReceivedAddress,
			transaction.Amount,
		)
	}
}

func callSendCoin(client services.WalletClient) {
	ctx := context.Background()
	req := services.SendCoinRequest{FromUserId: 1, ToUserId: 2, Amount: "0.00001"}
	res, err := client.SendCoin(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SendCoin Chainhash is %s", res.TransactionId)
}
