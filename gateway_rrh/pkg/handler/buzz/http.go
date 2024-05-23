package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type RoundRobinGrpcHandleListener struct {
	grpcServices []BuzzGrpc
	listRequests chan *HttpRequestWithResponse
	// Mutex
	m     sync.Mutex
	index int
}

type ReqParams struct {
	Str string `json:"str"`
}

type HttpRequestWithResponse struct {
	req *http.Request
	res http.ResponseWriter
}

func NewRoundRobinGrpcHandler(grpcServices []BuzzGrpc) *RoundRobinGrpcHandleListener {
	return &RoundRobinGrpcHandleListener{
		grpcServices: grpcServices,
		// Buffered channel
		listRequests: make(chan *HttpRequestWithResponse, len(grpcServices)),
		index:        0,
	}
}

func (r *RoundRobinGrpcHandleListener) HandleRequests(w http.ResponseWriter, rr *http.Request) {
	r.listRequests <- &HttpRequestWithResponse{req: rr, res: w}
}

func (r *RoundRobinGrpcHandleListener) Run() {
	for {
		r.processRequests()
	}
}

func (r *RoundRobinGrpcHandleListener) getNextGrpcClient() BuzzGrpc {
	r.m.Lock()
	defer r.m.Unlock()

	client := r.grpcServices[r.index]
	r.index = (r.index + 1) % len(r.grpcServices)

	return client
}

func (r *RoundRobinGrpcHandleListener) processRequests() {
	for reqWithRes := range r.listRequests {
		// Check if the request is a POST request
		if reqWithRes.req.Method != http.MethodPost {
			http.Error(reqWithRes.res, "Method not allowed", http.StatusMethodNotAllowed)

			return
		}

		var params ReqParams
		err := json.NewDecoder(reqWithRes.req.Body).Decode(&params)

		if err != nil {
			fmt.Println(err.Error())
			reqWithRes.res.WriteHeader(http.StatusBadRequest)

			return
		}

		grpcClient := r.getNextGrpcClient()
		str, err := grpcClient.Ping(params.Str)

		if err != nil {
			fmt.Println(err.Error())
			reqWithRes.res.WriteHeader(http.StatusForbidden)

			return
		}

		reqWithRes.res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(reqWithRes.res).Encode(str)

	}
}
