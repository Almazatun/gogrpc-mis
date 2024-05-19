package handler

import (
	"context"
	"time"

	buzz_proto "github.com/Almazatun/gogrpc-mis/gateway/pkg/genproto"
	"google.golang.org/grpc"
)

type FuzzGrpc interface {
	Ping(str string) (string, error)
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

func (fuzz *FuzzGrpcHandler) Ping(str string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := fuzz.client.Ping(ctx, &buzz_proto.PingRequest{
		Str: str,
	})

	if err != nil {
		return "", err
	}

	return res.Str, nil
}
