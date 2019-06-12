package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
)

func TestVerifyPlayerWithNoErr(t *testing.T) {
	gameTest := NewGame(false)
	rq := newReq("Join#A", "127.0.0.52")
	gameTest.Join(rq.From(), nil)
	err := verifyPlayer(gameTest, rq, messageSink)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPlayerWithErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	gameTest.Join("127.0.0.52", nil)
	rq := newReq("Auction#A", "127.0.0.52")
	gameTest.Join(rq.From(), nil)
	err := verifyPlayer(gameTest, rq, messageSink)
	if err == nil {
		t.Log(err)
		t.Fatal("Error was expected")
	}
}

func TestVerifyPhaseWithNoErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	rq := newReq("Join#A", "127.0.0.51")
	err := verifyPhase(gameTest, rq, messageSink)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPhaseWithErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	gameTest.phase = phase.End
	rq := newReq("Join#A", "127.0.0.51")
	err := verifyPhase(gameTest, rq, messageSink)
	if err == nil {
		t.Fatal("Error was expected")
	}
}
