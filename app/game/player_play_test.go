package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

var messageSink = func(p *player.Player, msg string) {}

func TestProcessRequestWithNoErr(t *testing.T) {
	gameTest := NewGame(false)
	rq := newReq("Join#A", "127.0.0.51")
	gameTest.Join(rq.From(), nil)
	err := processRequest(gameTest, rq, messageSink)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProcessRequestWithErr(t *testing.T) {
	gameTest := NewGame(false)
	rq := newReq("Card#A#B", "127.0.0.51")
	gameTest.Join(rq.From(), nil)
	err := processRequest(gameTest, rq, messageSink)
	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestProcessAuctionRequestWithNoErr(t *testing.T) {
	gameTest := NewGame(false)
	rq := newReq("Auction#75", "127.0.0.51")
	gameTest.Join(rq.From(), nil)
	err := processRequest(gameTest, rq, messageSink)
	if err != nil {
		t.Fatal(err)
	}
}
