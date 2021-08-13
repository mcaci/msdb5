package misc

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

var testPlayers Players

func init() {
	testPlayers = *NewPlayers(2)
	a := New(&Options{For2P: true, Name: "A"})
	a.Hand().Add(*card.MustID(34))
	testPlayers[0] = a
	b := New(&Options{For2P: true, Name: "B"})
	b.Hand().Add(*card.MustID(33))
	b.Hand().Add(*card.MustID(34))
	testPlayers[1] = b
}

func TestSuccessfulFindDataCorrespondsToA(t *testing.T) {
	i, err := testPlayers.Index(testPredicateA)
	if err != nil {
		t.Fatal(err)
	}
	if p := testPlayers[i]; !testPredicateA(p) {
		t.Fatalf("%s and %v are expected to be the same player", "A", p)
	}
}

func TestSuccessfulFindDataCorrespondsToB(t *testing.T) {
	testPredicateB := func(p Player) bool { return p.Name() == "B" }
	i, err := testPlayers.Index(testPredicateB)
	if err != nil {
		t.Fatal(err)
	}
	if p := testPlayers[i]; !testPredicateB(p) {
		t.Fatalf("%s and %v are expected to be the same player", "B", p)
	}
}
