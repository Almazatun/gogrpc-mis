package handler

import (
	"context"
	"time"

	buzz_proto "github.com/Almazatun/gogrpc-mis/gateway_rrh/pkg/genproto"
	"google.golang.org/grpc"
)

type BuzzGrpc interface {
	Ping(str string) (string, error)
}

type BuzzGrpcHandler struct {
	client buzz_proto.BuzzServiceClient
}

func NewBuzzGrpcHandler(client *grpc.ClientConn) *BuzzGrpcHandler {
	c := buzz_proto.NewBuzzServiceClient(client)

	return &BuzzGrpcHandler{
		client: c,
	}
}

func (buzz *BuzzGrpcHandler) Ping(str string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := buzz.client.Ping(ctx, &buzz_proto.PingRequest{
		Str: str,
	})

	if err != nil {
		return "", err
	}

	return res.Str, nil
}
