package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestPlayDoNoErr(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testObject := NewPlay("Play#1#Coin", "127.0.0.4", 0,
		testPlayer, playerset.Players{testPlayer}, board.New(), card.Coin)
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Play phase: %v", err)
	}
}

func TestPlayDoNoErrInRoundEnd(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6}
	testBoard := board.New()
	testBoard.PlayedCards().Add(testPlayedCards...)
	testObject := NewPlay("Play#1#Coin", "127.0.0.4", 0, testPlayer,
		playerset.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer},
		testBoard, card.Coin)
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Play phase: %v", err)
	}
}
