package team

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

var testPlayers Players

func init() {
	var a player.Player
	a.RegisterAs("A")
	testPlayers.Add(a)
	var b player.Player
	b.RegisterAs("B")
	b.Draw(func() card.ID { return 33 })
	testPlayers.Add(b)
}

func TestSuccessfulFindNoErr(t *testing.T) {
	if _, _, err := testPlayers.Find(func(p *player.Player) bool { return p.Has(33) }); err != nil {
		t.Fatal("Player not found with criteria p.Has(33)")
	}
}

func TestSuccessfulFindIndex(t *testing.T) {
	if index, _, _ := testPlayers.Find(func(p *player.Player) bool { return p.Has(33) }); index != 1 {
		t.Fatal("Player not found with criteria p.Has(33)")
	}
}

func TestSuccessfulFindDataCorresponds(t *testing.T) {
	isPlayerACheck := func(p *player.Player) bool { return p.Name() == "A" }
	if _, player, _ := testPlayers.Find(isPlayerACheck); !isPlayerACheck(player) {
		t.Fatalf("%s and %v are expected to be the same player", "A", player)
	}
}

func TestUnsuccessfulFind(t *testing.T) {
	if _, _, err := testPlayers.Find(func(p *player.Player) bool { return p.Has(24) }); err == nil {
		t.Fatal(err)
	}
}
