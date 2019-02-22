package board

import (
	"testing"
)

func TestPlayerPlaysOneCardAndCardIsOnTheBoard(t *testing.T) {
	b := New()
	b.Join("A", "100.0.0.1")
	b.Play("1", "Coin", "100.0.0.1")
	if !b.PlayedCards().Has(1) {
		t.Fatal("Board should have One of Coin as played card")
	}
}

func TestWhenPlay5CardsTriggerMove(t *testing.T) {
	b := New()
	b.Join("A", "100.0.0.1")
	b.Join("B", "100.0.0.2")
	b.Join("C", "100.0.0.3")
	b.Join("D", "100.0.0.4")
	b.Join("E", "100.0.0.5")
	b.Play("5", "Coin", "100.0.0.1")
	b.Play("3", "Cup", "100.0.0.2")
	b.Play("1", "Coin", "100.0.0.3")
	b.Play("4", "Sword", "100.0.0.4")
	b.Play("8", "Cudgel", "100.0.0.5")
	if len(*b.PlayedCards()) != 0 {
		t.Fatal("Board should have triggered the move to winning player")
	}
}
