package servers

import (
	"context"

	"github.com/btcsuite/btcd/rpcclient"
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

	res.Balance = decimal.NewFromFloat(float64(amount) / 100000000).StringFixed(8)

	return &res, nil
}

func (s *Wallet) getAccount(user_id uint64) string {
	if user_id == 1 {
		return "client1"
	}
	return "tno201806"
}
