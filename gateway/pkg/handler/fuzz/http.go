package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	handler "github.com/Almazatun/gogrpc-mis/gateway/pkg/handler/buzz"
)

type FuzzHttpHandler struct {
	grpc FuzzGrpc
}

type FuzzHttp interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

func NewFuzzHttpHandler(grpc FuzzGrpc) FuzzHttp {
	return &FuzzHttpHandler{
		grpc: grpc,
	}
}

func (h *FuzzHttpHandler) Ping(w http.ResponseWriter, r *http.Request) {
	var params handler.ReqParams
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	str, err := h.grpc.Ping(params.Str)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusForbidden)

		return
	}

	json.NewEncoder(w).Encode(str)
}
