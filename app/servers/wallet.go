package servers

import (
	"context"
	"github.com/tnakade/tno_exercise/app/proto/services"
	"math/rand"
	"time"
	"github.com/shopspring/decimal"
)

type Wallet struct {
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (s *Wallet) GetBalance(ctx context.Context, msg *services.GetBalanceRequest) (*services.GetBalanceResponse, error) {
	res := services.GetBalanceResponse{}
	rand.Seed(time.Now().UnixNano())
	amount := decimal.NewFromFloat(float64(rand.Intn(10000000)) / 1000000)
	res.Balance = amount.StringFixed(8)

	return &res, nil
}
