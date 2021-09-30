package briscola_test

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v3/briscola"
)

var testPlayers briscola.Players

func init() {
	testPlayers = *briscola.NewPlayers(2)
	a := briscola.NewPlayer("A")
	a.Hand().Add(*card.MustID(34))
	testPlayers[0] = a
	b := briscola.NewPlayer("B")
	b.Hand().Add(*card.MustID(33))
	b.Hand().Add(*card.MustID(34))
	testPlayers[1] = b
}

func testPredicateA(p briscola.Player) bool { return p.Name() == "A" }

func TestSuccessfulFindDataCorrespondsToA(t *testing.T) {
	t.Parallel()
	i, err := testPlayers.Index(testPredicateA)
	if err != nil {
		t.Fatal(err)
	}
	if p := testPlayers[i]; !testPredicateA(*p) {
		t.Fatalf("%s and %v are expected to be the same player", "A", p)
	}
}

func TestSuccessfulFindDataCorrespondsToB(t *testing.T) {
	t.Parallel()
	testPredicateB := func(p briscola.Player) bool { return p.Name() == "B" }
	i, err := testPlayers.Index(testPredicateB)
	if err != nil {
		t.Fatal(err)
	}
	if p := testPlayers[i]; !testPredicateB(*p) {
		t.Fatalf("%s and %v are expected to be the same player", "B", p)
	}
}
