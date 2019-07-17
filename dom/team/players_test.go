package team

import (
	"testing"

	"github.com/mcaci/msdb5/dom/player"
)

var testPlayers Players

func init() {
	var a player.Player
	a.RegisterAs("A")
	testPlayers.Add(&a)
	var b player.Player
	b.RegisterAs("B")
	b.Hand().Add(33)
	testPlayers.Add(&b)
}

func TestSuccessfulFindNoErr(t *testing.T) {
	if _, p := testPlayers.Find(func(p *player.Player) bool { return p.Has(33) }); p == nil {
		t.Fatal("Player not found with criteria p.Has(33)")
	}
}

func TestSuccessfulFindIndex(t *testing.T) {
	if index, _ := testPlayers.Find(func(p *player.Player) bool { return p.Has(33) }); index != 1 {
		t.Fatal("Player not found with criteria p.Has(33)")
	}
}

func TestSuccessfulFindDataCorresponds(t *testing.T) {
	isPlayerACheck := func(p *player.Player) bool { return p.Name() == "A" }
	if _, player := testPlayers.Find(isPlayerACheck); !isPlayerACheck(player) {
		t.Fatalf("%s and %v are expected to be the same player", "A", player)
	}
}

func TestUnsuccessfulFind(t *testing.T) {
	if _, p := testPlayers.Find(func(p *player.Player) bool { return p.Has(24) }); p != nil {
		t.Fatal("Player should not be found")
	}
}
