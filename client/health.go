package client

import (
	"context"
	"time"

	"github.com/TomasCruz/grpc-web-fahrenheit/api"
)

// Health verifies client is online
func (g grpcClient) Health() (status bool, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	response, err := g.c.Health(ctx, &api.NoParamsMsg{})
	if err != nil {
		return
	}

	status = response.GetHealth()
	return
}
