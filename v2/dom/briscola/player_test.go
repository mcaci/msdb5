package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestJoinPlayerName(t *testing.T) {
	if p := NewB2Player("Me"); p.Name() != "Me" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerPileIsEmpty(t *testing.T) {
	if p := NewB2Player(""); len(*p.Pile()) != 0 {
		t.Fatal("Pile should be empty")
	}
}

func TestPlayerDrawsOneCard(t *testing.T) {
	p := NewB2Player("")
	p.Hand().Add(*card.MustID(1))
	if len(*p.Hand()) != 1 {
		t.Fatalf("Expecting player to have drawn %v", *card.MustID(1))
	}
}
