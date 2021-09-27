package briscola_test

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v3/briscola"
)

func TestJoinPlayerName(t *testing.T) {
	if p := briscola.NewPlayer("Me"); p.Name() != "Me" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerPileIsEmpty(t *testing.T) {
	if p := briscola.NewPlayer(""); len(*p.Pile()) != 0 {
		t.Fatal("Pile should be empty")
	}
}

func TestPlayerDrawsOneCard(t *testing.T) {
	p := briscola.NewPlayer("")
	p.Hand().Add(*card.MustID(1))
	if len(*p.Hand()) != 1 {
		t.Fatalf("Expecting player to have drawn %v", *card.MustID(1))
	}
}
