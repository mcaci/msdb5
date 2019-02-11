package board

import (
	"strconv"
	"testing"
)

func TestCardNominatedWithOriginInfo(t *testing.T) {
	b := New()
	b.Nominate("1", "Coin", "100.1.1.1")
	if *b.NominatedCard() != 1 {
		t.Fatal("Nominated card is not 1 of Coin")
	}
}

func TestCardNominatedWithOriginInfo_SetNominatedPlayerA(t *testing.T) {
	b := New()
	b.Join("A", "100.1.1.1")
	b.Join("B", "100.1.1.2")
	card := (*b.Players()[0].Hand())[0]
	b.Nominate(strconv.Itoa((int)(card.Number())), card.Seed().String(), "100.1.1.2")
	if b.NominatedPlayer().Name() != "A" {
		t.Fatal("Nominated player is not A")
	}
}

func TestCardNominatedWithOriginInfo_SetNominatedPlayerB(t *testing.T) {
	b := New()
	b.Join("A", "100.1.1.1")
	b.Join("B", "100.1.1.2")
	card := (*b.Players()[1].Hand())[0]
	b.Nominate(strconv.Itoa((int)(card.Number())), card.Seed().String(), "100.1.1.1")
	if b.NominatedPlayer().Name() != "B" {
		t.Fatal("Nominated player is not B")
	}
}
