package client

import (
	"fmt"

	"github.com/TomasCruz/grpc-web-fahrenheit/api"
	"github.com/TomasCruz/grpc-web-fahrenheit/model"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type grpcClient struct {
	conn *grpc.ClientConn
	c    api.ConvertorClient
}

// InitializeClient establishes connection to gRCP server
func InitializeClient(host, port string) (client model.Client, err error) {
	g := grpcClient{}
	if g.conn, err = grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure()); err != nil {
		err = errors.WithStack(err)
		return
	}

	g.c = api.NewConvertorClient(g.conn)
	client = g

	return
}
