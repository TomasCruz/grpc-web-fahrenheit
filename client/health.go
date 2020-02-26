package client

import (
	"context"
	"time"

	"github.com/TomasCruz/grpc-web-fahrenheit/model"
	"github.com/pkg/errors"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Health verifies client is online
func (g grpcClient) Health() (status bool, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	response, err := g.h.Check(ctx, &healthpb.HealthCheckRequest{})
	if err != nil {
		err = errors.Wrap(model.ErrClient, err.Error())
		return
	}

	status = response.GetStatus() == healthpb.HealthCheckResponse_SERVING
	return
}
