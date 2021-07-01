package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/player"
)

var testPlayers Players

func init() {
	a := player.New(&player.Options{For2P: true, Name: "A"})
	a.Hand().Add(*card.MustID(34))
	testPlayers.Add(a)
	b := player.New(&player.Options{For2P: true, Name: "B"})
	b.Hand().Add(*card.MustID(33))
	b.Hand().Add(*card.MustID(34))
	testPlayers.Add(b)
}

func TestSuccessfulFindDataCorresponds(t *testing.T) {
	i, err := testPlayers.SelectIndex(testPredicate)
	if err != nil {
		t.Fatal(err)
	}
	if p := testPlayers[i]; !testPredicate(p) {
		t.Fatalf("%s and %v are expected to be the same player", "A", p)
	}
}
