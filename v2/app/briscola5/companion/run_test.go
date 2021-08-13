package companion

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

func TestRunMinimalExample(t *testing.T) {
	in := companionIn{}
	in.Players = *briscola.NewPlayers(5)
	in.Player = in.Players[0]
	out := Run(in)
	if out.Companion != 0 {
		t.Errorf("Not possible that a companion was found: %d", out.Companion)
	}
	var card briscola.Card
	if out.Briscola != card {
		t.Errorf("Not possible that a card was found: %v", out.Briscola)
	}
}

func TestRunExampleNominal(t *testing.T) {
	in := companionIn{}
	in.Players = *briscola.NewPlayers(5)
	in.Players[0].Hand().Add(*card.MustID(1), *card.MustID(1))
	in.Players[1].Hand().Add(*card.MustID(11), *card.MustID(23))
	in.Players[2].Hand().Add(*card.MustID(11), *card.MustID(22))
	in.Players[3].Hand().Add(*card.MustID(21), *card.MustID(21))
	in.Players[4].Hand().Add(*card.MustID(31), *card.MustID(31))
	in.Player = in.Players[3]
	out := Run(in)
	if out.Companion != 1 {
		t.Errorf("Not possible that a companion was found: %d", out.Companion)
	}
	if out.Briscola.ToID() != 23 {
		t.Errorf("Not possible that a card was found: %v", out.Briscola)
	}
}

func TestRunExampleAnother(t *testing.T) {
	in := companionIn{}
	in.Players = *briscola.NewPlayers(5)
	in.Players[0].Hand().Add(*card.MustID(1), *card.MustID(1), *card.MustID(1))
	in.Players[1].Hand().Add(*card.MustID(11), *card.MustID(22), *card.MustID(11))
	in.Players[2].Hand().Add(*card.MustID(23), *card.MustID(29), *card.MustID(11))
	in.Players[3].Hand().Add(*card.MustID(21), *card.MustID(21), *card.MustID(11))
	in.Players[4].Hand().Add(*card.MustID(31), *card.MustID(31), *card.MustID(31))
	in.Player = in.Players[3]
	out := Run(in)
	if out.Companion != 2 {
		t.Errorf("Not possible that a companion was found: %d", out.Companion)
	}
	if out.Briscola.ToID() != 23 {
		t.Errorf("Not possible that a card was found: %v", out.Briscola)
	}
}
