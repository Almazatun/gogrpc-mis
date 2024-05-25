package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type RoundRobinGrpcHandleListener struct {
	grpcServices []BuzzGrpc
	listRequests chan HttpReq
	replay       chan string
	err          chan error
	// Mutex
	m     sync.Mutex
	index int
}

type ReqParams struct {
	Str string `json:"str"`
}

type HttpReq struct {
	params ReqParams
}

func NewRoundRobinGrpcHandler(grpcServices []BuzzGrpc) *RoundRobinGrpcHandleListener {
	return &RoundRobinGrpcHandleListener{
		grpcServices: grpcServices,
		// Buffered channel
		listRequests: make(chan HttpReq, len(grpcServices)),
		// Unbuffered channel
		replay: make(chan string),
		err:    make(chan error),
		index:  0,
	}
}

func (r *RoundRobinGrpcHandleListener) HandleRequests(w http.ResponseWriter, rr *http.Request) {
	// Check if the request is a POST request
	if rr.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	var params ReqParams
	err := json.NewDecoder(rr.Body).Decode(&params)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	r.listRequests <- HttpReq{params: params}

	// fmt.Println("Request received")

	if msg := r.listenReplayMsg(); msg != "" {
		// fmt.Println("Request processed", time.Now())

		json.NewEncoder(w).Encode(msg)
		return
	}

	if err := r.listenErr(); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusForbidden)

		return
	}

}

func (r *RoundRobinGrpcHandleListener) Run() {
	for {
		r.processRequests()
	}
}
func (p *RoundRobinGrpcHandleListener) listenReplayMsg() string {
	for err := range p.replay {
		return err
	}

	return ""
}

func (p *RoundRobinGrpcHandleListener) listenErr() error {
	for err := range p.err {
		return err
	}

	return nil
}

func (r *RoundRobinGrpcHandleListener) getNextGrpcClient() BuzzGrpc {
	r.m.Lock()
	defer r.m.Unlock()

	client := r.grpcServices[r.index]
	r.index = (r.index + 1) % len(r.grpcServices)

	// fmt.Println(r.index, "Index")

	return client
}

func (r *RoundRobinGrpcHandleListener) processRequests() {
	for reqWithRes := range r.listRequests {
		grpcClient := r.getNextGrpcClient()

		str, err := grpcClient.Ping(reqWithRes.params.Str)

		if err != nil {
			fmt.Println(err.Error())

			r.err <- err
			continue
		}

		r.replay <- str
	}
}
