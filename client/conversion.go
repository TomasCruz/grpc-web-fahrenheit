package client

import (
	"context"
	"time"

	"github.com/TomasCruz/grpc-web-fahrenheit/api"
)

func (g grpcClient) C2F(c float64) (f float64, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	response, err := g.c.C2F(ctx, &api.ConversionMsg{Number: c})
	if err != nil {
		return
	}

	f = response.GetNumber()
	return
}

func (g grpcClient) F2C(f float64) (c float64, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	response, err := g.c.F2C(ctx, &api.ConversionMsg{Number: f})
	if err != nil {
		return
	}

	c = response.GetNumber()
	return
}
