package server

import (
	"context"

	"github.com/arthurcgc/grpc-learning/protos/currency"
	"github.com/sirupsen/logrus"
)

type Currency struct {
	currency.UnimplementedCurrencyServer
	logger *logrus.Logger
}

func New() *Currency {
	return &Currency{logger: logrus.New()}
}

func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.logger.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &currency.RateResponse{Rate: 0.5}, nil
}
