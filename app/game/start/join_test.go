package start

import (
	"testing"

	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type fakeGame team.Players

func (f fakeGame) Players() team.Players {
	return team.Players(f)
}

func TestRegisterPlayerHasLocalhostOrigin(t *testing.T) {
	p := player.New()
	gameTest := fakeGame{p}
	playerInfo := "localhost"
	Join(gameTest, playerInfo, make(chan []byte))
	if p == nil {
		t.Fatalf("Player %s is expected to exist", playerInfo)
	}
}
