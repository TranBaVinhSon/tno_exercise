package servers

import (
	"context"

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

	res.Balance = decimal.NewFromFloat(float64(amount) / 100000000).StringFixed(8)

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

	res.ChainHash = hash.String()

	return &res, nil
}

func (s *Wallet) getAccount(user_id uint64) string {
	if user_id == 1 {
		return "client1"
	}
	return "tno201806"
}
