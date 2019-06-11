package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

var messageSink = func(p *player.Player, msg string) {}

func TestProcessRequestWithNoErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	err := processRequest(gameTest, "Join#A", "127.0.0.51", messageSink)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProcessRequestWithErr(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", nil)
	err := processRequest(gameTest, "Card#A#B", "127.0.0.51", messageSink)
	if err == nil {
		t.Fatal("Error was expected")
	}
}
