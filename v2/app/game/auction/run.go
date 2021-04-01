package auction

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

func Run(players briscola5.Players, listenFor func(context.Context, func())) struct {
	Score  briscola5.AuctionScore
	Caller uint8
} {
	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- 60 + rand.Intn(60) })
	go func() {
		<-done
		cancel()
		close(numbers)
	}()

	var score briscola5.AuctionScore
	var currID uint8

	for n := range numbers {
		r := Round(struct {
			curr, prop briscola5.AuctionScore
			currID     uint8
			players    briscola5.Players
			cmpF       func(briscola5.AuctionScore, briscola5.AuctionScore) int8
		}{
			curr:    score,
			prop:    briscola5.AuctionScore(n),
			currID:  currID,
			players: players,
			cmpF:    callCmp,
		})
		score = r.s
		currID = r.id
		if !r.end {
			continue
		}
		done <- struct{}{}
		close(done)
	}
	return struct {
		Score  briscola5.AuctionScore
		Caller uint8
	}{
		Score:  score,
		Caller: players.MustIndex(notFolded),
	}
}

func notFolded(p *briscola5.Player) bool { return !briscola5.Folded(p) }

func callCmp(curr, prop briscola5.AuctionScore) int8 {
	var jsonReq = fmt.Sprintf(`{"current":%d,"proposed":%d}`, curr, prop)
	res, err := http.Post("http://localhost:8080/cmp", "application/json", strings.NewReader(jsonReq))
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	var rs struct {
		Cmp int8 `json:"cmp"`
	}
	rserr := json.NewDecoder(res.Body).Decode(&rs)
	if rserr != nil {
		log.Println(err)
	}
	return rs.Cmp
}
