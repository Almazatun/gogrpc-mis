package handler

import (
	"context"

	buzz_proto "github.com/Almazatun/gogrpc-mis/service_buzz/pkg/genproto"
	"github.com/Almazatun/gogrpc-mis/service_buzz/pkg/types"
	"google.golang.org/grpc"
)

type BuzzGrpcHandler struct {
	service types.BuzzService
	buzz_proto.UnimplementedBuzzServiceServer
}

func NewBuzzGrpcHandler(grpc *grpc.Server, service types.BuzzService) {
	buzzGrpcHandler := &BuzzGrpcHandler{
		service: service,
	}
	// register service
	buzz_proto.RegisterBuzzServiceServer(grpc, buzzGrpcHandler)
}

func (h *BuzzGrpcHandler) Ping(ctx context.Context, req *buzz_proto.PingRequest) (*buzz_proto.PongResponse, error) {
	res := h.service.Ping()

	return &buzz_proto.PongResponse{
		Str: res,
	}, nil
}
