package handler

import (
	"context"
	"time"

	buzz_proto "github.com/Almazatun/gogrpc-mis/gateway/pkg/genproto"
	"google.golang.org/grpc"
)

type FuzzGrpc interface {
	Ping(str string) (*string, error)
}

type FuzzGrpcHandler struct {
	client buzz_proto.FuzzServiceClient
}

func NewFuzzGrpcHandler(client *grpc.ClientConn) *FuzzGrpcHandler {
	c := buzz_proto.NewFuzzServiceClient(client)

	return &FuzzGrpcHandler{
		client: c,
	}
}

func (buzz *FuzzGrpcHandler) Ping(str string) (*buzz_proto.PongResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := buzz.client.Ping(ctx, &buzz_proto.PingRequest{
		Str: str,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
