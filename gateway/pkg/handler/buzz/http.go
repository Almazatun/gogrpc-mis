package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BuzzHttpHandler struct {
	grpc BuzzGrpc
}

type BuzzHttp interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type ReqParams struct {
	Str string `json:"str"`
}

func NewBuzzHttpHandler(grpc BuzzGrpc) BuzzHttp {
	return &BuzzHttpHandler{
		grpc: grpc,
	}
}

func (h *BuzzHttpHandler) Ping(w http.ResponseWriter, r *http.Request) {
	var params ReqParams
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
