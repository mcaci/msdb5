package board

import "testing"

func TestWhenPlayLessThan5CardsRoundIsNotOver(t *testing.T) {
	b := New()
	isRoundEnded := b.PlayedCardIs(5)
	if isRoundEnded {
		t.Fatal("Round is not ended")
	}
}

func TestWhenPlay5CardsRoundIsOver(t *testing.T) {
	b := New()
	isRoundEnded := b.PlayedCardIs(5)
	isRoundEnded = b.PlayedCardIs(13)
	isRoundEnded = b.PlayedCardIs(1)
	isRoundEnded = b.PlayedCardIs(24)
	isRoundEnded = b.PlayedCardIs(38)
	if !isRoundEnded {
		t.Fatal("Round should be over")
	}
}
