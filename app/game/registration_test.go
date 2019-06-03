package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func TestRegisterPlayerHasLocalhostOrigin(t *testing.T) {
	testGame := NewGame(false)
	playerInfo := "localhost"
	testGame.Join(playerInfo, make(chan []byte))
	testPlayers := testGame.players
	_, _, err := testPlayers.Find(func(p *player.Player) bool { return p.IsSameHost(playerInfo) })
	if err != nil {
		t.Fatalf("Player %s is expected to exist", playerInfo)
	}
}
