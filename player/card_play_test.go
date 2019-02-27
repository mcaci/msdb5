package player

import (
	"testing"
)

func TestPlayerPlaysCard(t *testing.T) {
	p := New()
	p.Hand().Add(1)
	oldHand := *p.Hand()
	card, err := p.Play("1", "Coin")
	if err != nil {
		t.Logf("Card played: %v", card)
		t.Logf("Hand before playing: %v", oldHand)
		t.Fatal("Card should come from player's hand")
	}
}

func TestHandSizeChangesIfPlayerPlaysCardInHand(t *testing.T) {
	p := New()
	p.Hand().Add(1)
	oldHand := *p.Hand()
	p.Play("1", "Coin")
	if len(oldHand) != len(*p.Hand())+1 {
		t.Fatal("In case of error handsize should not change")
	}
}

func TestErrIfPlayerPlaysCardNotInHand(t *testing.T) {
	p := New()
	p.Hand().Add(1)
	_, err := p.Play("2", "Coin")
	if err == nil {
		t.Fatal("Card should come from player's hand")
	}
}

func TestNoCardReturnedIfPlayerPlaysCardNotInHand(t *testing.T) {
	p := New()
	p.Hand().Add(1)
	card, _ := p.Play("2", "Coin")
	if card != 0 {
		t.Fatal("Card should come from player's hand")
	}
}

func TestHandSizeDoesntChangeIfPlayerPlaysCardNotInHand(t *testing.T) {
	p := New()
	p.Hand().Add(1)
	oldHand := *p.Hand()
	p.Play("2", "Coin")
	if len(oldHand) != len(*p.Hand()) {
		t.Fatal("In case of error handsize should not change")
	}
}
