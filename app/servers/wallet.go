package servers

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/shopspring/decimal"
	"github.com/tnakade/tno_exercise/app/proto/services"
)

type Wallet struct {
	client *rpcclient.Client
}

func NewWallet(client *rpcclient.Client) *Wallet {
	wallet := &Wallet{}

	wallet.client = client

	return wallet
}

func (s *Wallet) GetBalance(ctx context.Context, msg *services.GetBalanceRequest) (*services.GetBalanceResponse, error) {
	res := services.GetBalanceResponse{}

	amount, err := s.client.GetBalance(s.getAccount(msg.UserId))
	if err != nil {
		return &res, err
	}

	res.Balance = s.makeBTCString(int64(amount))

	return &res, nil
}

func (s *Wallet) SendCoin(ctx context.Context, msg *services.SendCoinRequest) (*services.SendCoinResponse, error) {
	res := services.SendCoinResponse{}

	amount_decimal, err := decimal.NewFromString(msg.Amount)
	if err != nil {
		return &res, err
	}

	amount_float64, _ := amount_decimal.Float64()

	amount, err := btcutil.NewAmount(amount_float64)
	if err != nil {
		return &res, err
	}

	addresses, err := s.client.GetAddressesByAccount(s.getAccount(msg.ToUserId))
	if err != nil {
		return &res, err
	}

	hash, err := s.client.SendFrom(s.getAccount(msg.FromUserId), addresses[0], amount)
	if err != nil {
		return &res, err
	}

	res.TransactionId = hash.String()

	return &res, nil
}

func (s *Wallet) GetTransactions(ctx context.Context, msg *services.GetTransactionsRequest) (*services.GetTransactionsResponse, error) {
	res := services.GetTransactionsResponse{}

	transactions_result, err := s.client.ListTransactions(s.getAccount(msg.UserId))
	if err != nil {
		return &res, err
	}

	account_indexed_by_address := s.getAccountIndexedByAddress()
	transactions := make([]*services.Transaction, len(transactions_result))

	for i, transaction := range transactions_result {
		account := account_indexed_by_address[transaction.Address]
		user_id := s.getUserId(account)

		transactions[i] = &services.Transaction{
			Id:        transaction.TxID,
			Category:  transaction.Category,
			Abandoned: strconv.FormatBool(transaction.Abandoned),
			ReceivedAddress: &services.TransactionReceivedAddress{
				Id: transaction.Address,
				User: &services.User{
					Id:      fmt.Sprintf("%d", user_id),
					Name:    s.getUserName(user_id),
					Account: account,
				},
			},
			Amount:     fmt.Sprintf("%f", math.Abs(transaction.Amount)),
			SendAt:     time.Unix(transaction.Time, 0).String(),
			ReceivedAt: time.Unix(transaction.TimeReceived, 0).String(),
		}
	}

	res.Transactions = transactions

	return &res, nil
}

func (s *Wallet) getAccount(user_id uint64) string {
	if user_id == 1 {
		return "client1"
	}
	return "tno201806"
}

func (s *Wallet) getUserId(account string) uint64 {
	if account == "client1" {
		return 1
	}
	return 2
}

func (s *Wallet) getUserName(user_id uint64) string {
	if user_id == 1 {
		return "玉井"
	}
	return "中出商会"
}

func (s *Wallet) getAccountIndexedByAddress() map[string]string {
	account_indexed_by_address := make(map[string]string)

	amount_indexed_by_account, err := s.client.ListAccounts()
	if err != nil {
		return account_indexed_by_address
	}

	for account, _ := range amount_indexed_by_account {
		addresses, err := s.client.GetAddressesByAccount(account)
		if err != nil {
			continue
		}

		for _, address := range addresses {
			account_indexed_by_address[address.String()] = account
		}
	}

	return account_indexed_by_address
}

func (s *Wallet) makeBTCString(satoshi_amount int64) string {
	return decimal.NewFromFloat(float64(satoshi_amount) / 100000000).StringFixed(8)
}
