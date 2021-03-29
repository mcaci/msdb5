package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/player"
)

var testPlayers Players

func init() {
	var a player.Player
	a.RegisterAs("A")
	a.Hand().Add(*card.MustID(34))
	testPlayers.Add(&a)
	var b player.Player
	b.RegisterAs("B")
	b.Hand().Add(*card.MustID(33))
	b.Hand().Add(*card.MustID(34))
	testPlayers.Add(&b)
}

func TestSuccessfulFindDataCorresponds(t *testing.T) {
	isPlayerACheck := func(p *player.Player) bool { return p.Name() == "A" }
	if p := testPlayers[testPlayers.MustIndex(isPlayerACheck)]; !isPlayerACheck(p) {
		t.Fatalf("%s and %v are expected to be the same player", "A", p)
	}
}
