package exchange

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func TestExchangeDoNoErr(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6, 7}
	testBoard := board.New()
	testBoard.SideDeck().Add(testPlayedCards...)
	testObject := NewExchangeCards("Exchange#1#Coin", "127.0.0.2", testBoard.SideDeck())
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatal("Unexpected error when exchanging cards phase")
	}
}

func TestExchangeOneCardPicksFirst(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6, 7}
	testBoard := board.New()
	testBoard.SideDeck().Add(testPlayedCards...)
	testObject := NewExchangeCards("Exchange#1#Coin", "127.0.0.2", testBoard.SideDeck())
	testObject.Do(testPlayer)
	if !testPlayer.Hand().Has(2) {
		t.Fatalf("Cards were not exchanged properly. Current hand: %v", testPlayer.Hand())
	}
}

func TestExchangeOneCardAndSideDeckSizeDoesNotChange(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6, 7}
	testBoard := board.New()
	testBoard.SideDeck().Add(testPlayedCards...)
	testObject := NewExchangeCards("Exchange#1#Coin", "127.0.0.2", testBoard.SideDeck())
	testObject.Do(testPlayer)
	if len(testPlayedCards) != len(*testBoard.SideDeck()) {
		t.Fatalf("Cards were not exchanged properly. Current hand: %v", testBoard.SideDeck())
	}
}

func TestExchangeOneCardAndHandSizeDoesNotChange(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	handLengthBefore := len(*testPlayer.Hand())
	testPlayedCards := deck.Cards{2, 3, 4, 6, 7}
	testBoard := board.New()
	testBoard.SideDeck().Add(testPlayedCards...)
	testObject := NewExchangeCards("Exchange#1#Coin", "127.0.0.2", testBoard.SideDeck())
	testObject.Do(testPlayer)
	if handLengthBefore != len(*testPlayer.Hand()) {
		t.Fatalf("Cards were not exchanged properly. Current hand: %v", testPlayer.Hand())
	}
}

func TestExchangeEndPhaseDoesNothingAndReturnsNoErr(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6, 7}
	testBoard := board.New()
	testBoard.SideDeck().Add(testPlayedCards...)
	testObject := NewExchangeCards("Exchange#0", "127.0.0.2", testBoard.SideDeck())
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatal("Unexpected error when exchanging cards phase")
	}
}
