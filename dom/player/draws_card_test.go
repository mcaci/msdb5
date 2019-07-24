package player

import (
	"testing"
)

func TestPlayerHasNoCardsAtStartGame(t *testing.T) {
	if p := New(); !IsHandEmpty(p) {
		t.Fatal("Player should not have cards at creation")
	}
}

func TestPlayerDrawsOneCard(t *testing.T) {
	p := New()
	p.Hand().Add(1)
	plPredicate := IsCardInHand(1)
	if !plPredicate(p) {
		t.Fatalf("Expecting player to have drawn %v", 1)
	}
}
