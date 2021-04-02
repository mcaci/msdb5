package srv

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

type aucreq struct {
	Current  briscola5.AuctionScore `json:"current"`
	Proposed briscola5.AuctionScore `json:"proposed"`
}

func Cmp(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("/cmp executing")
	var req aucreq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	type cmpres struct {
		Cmp int8 `json:"cmp"`
	}
	res := make(chan cmpres)
	go func(chan<- cmpres) {
		res <- cmpres{Cmp: int8(briscola5.Cmp(req.Current, req.Proposed))}
	}(res)
	select {
	case <-ctx.Done():
		log.Println("/cmp exiting with error", ctx.Err())
		rw.Write([]byte(ctx.Err().Error()))
	case r := <-res:
		log.Println("/cmp exiting", r)
		json.NewEncoder(rw).Encode(&r)
	}
}

func CmpAndSet(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("/cmpandset executing")
	var req aucreq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	type setres struct {
		Result briscola5.AuctionScore `json:"result"`
	}
	res := make(chan setres)
	go func(chan<- setres) {
		res <- setres{Result: briscola5.CmpAndSet(req.Current, req.Proposed)}
	}(res)
	select {
	case <-ctx.Done():
		log.Println("/cmpandset exiting with error", ctx.Err())
		rw.Write([]byte(ctx.Err().Error()))
	case r := <-res:
		log.Println("/cmpandset exiting", r)
		json.NewEncoder(rw).Encode(&r)
	}
}
