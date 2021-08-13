package auction

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola/player"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

func Run(auctIn struct {
	Players briscola.Players
	CmpF    func(briscola5.AuctionScore, briscola5.AuctionScore) int8
}) struct {
	Score  briscola5.AuctionScore
	Caller uint8
} {
	if auctIn.CmpF == nil {
		log.Println("using default auction score comparer")
		auctIn.CmpF = callCmp
	}

	var score briscola5.AuctionScore
	var currID uint8
	for {
		r := Round(struct {
			curr, prop briscola5.AuctionScore
			currID     uint8
			players    briscola.Players
			cmpF       func(briscola5.AuctionScore, briscola5.AuctionScore) int8
		}{
			curr:    score,
			prop:    briscola5.AuctionScore(60 + rand.Intn(60)),
			currID:  currID,
			players: auctIn.Players,
			cmpF:    auctIn.CmpF,
		})
		score = r.s
		currID = r.id
		if !r.end {
			continue
		}
		caller, err := auctIn.Players.Index(player.NotFolded)
		if err != nil {
			log.Println(err)
			continue
		}
		return struct {
			Score  briscola5.AuctionScore
			Caller uint8
		}{
			Score:  score,
			Caller: caller,
		}
	}
}

func callCmp(curr, prop briscola5.AuctionScore) int8 {
	var jsonReq = fmt.Sprintf(`{"current":%d,"proposed":%d}`, curr, prop)
	res, err := http.Post("http://localhost:8082/cmp", "application/json", strings.NewReader(jsonReq))
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
