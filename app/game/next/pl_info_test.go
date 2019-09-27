package next

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func testPlayers() team.Players {
	p := player.New()
	p.Join("127.0.0.1")
	p.RegisterAs("A")
	return team.Players{p}
}

func TestNextPlayer(t *testing.T) {
	testObj := NewPlInfo(phase.PlayingCards, testPlayers(), *card.MustID(1), &set.Cards{*card.MustID(2)}, "127.0.0.1")
	next := Player(testObj)
	if next.Name() != "A" {
		t.Fatal(next)
	}
}
