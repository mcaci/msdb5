package deck

import (
	"testing"
)

func TestCreateSet(t *testing.T) {
	cards := Cards{15}
	if !cards.Has(15) {
		t.Fatalf("There should be the 5 of Cup card in the set")
	}
}

func TestAddCardToSet(t *testing.T) {
	cards := Cards{}
	cards.Add(33)
	if !cards.Has(33) {
		t.Fatal("There should be the 3 of Cudgel card in the set")
	}
}

func TestRemoveCardFromSet(t *testing.T) {
	cards := Cards{15}
	cards.Remove(0)
	if cards.Has(15) {
		t.Fatalf("Cards should be empty")
	}
}

func TestMovedCardsAreAddedToDestination(t *testing.T) {
	playedCards := Cards{2, 3, 4, 5, 6}
	playerPile := Cards{}
	playerPile.Add(playedCards...)
	playedCards.Clear()
	if len(playerPile) == 0 {
		t.Fatalf("Cards did not move to player pile")
	}
}

func TestMovedCardsAreNotInOriginalSet(t *testing.T) {
	playedCards := Cards{2, 3, 4, 5, 6}
	playerPile := Cards{}
	playerPile.Add(playedCards...)
	playedCards.Clear()
	if len(playedCards) != 0 {
		t.Fatalf("Cards were not moved from played cards")
	}
}
