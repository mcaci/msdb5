package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
)

func TestVerifyPlayerWithNoErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	err := verifyPlayer(gameTest, "Join#A", "127.0.0.51", messageSink)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPlayerWithErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	gameTest.Join("127.0.0.52", nil)
	err := verifyPlayer(gameTest, "Auction#A", "127.0.0.52", messageSink)
	if err == nil {
		t.Log(err)
		t.Fatal("Error was expected")
	}
}

func TestVerifyPhaseWithNoErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	err := verifyPhase(gameTest, "Join#A", "127.0.0.51", messageSink)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPhaseWithErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	gameTest.phase = phase.End
	err := verifyPhase(gameTest, "Join#A", "127.0.0.51", messageSink)
	if err == nil {
		t.Fatal("Error was expected")
	}
}
