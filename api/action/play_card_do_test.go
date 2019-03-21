package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestPlayDoNoErr(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testObject := NewPlay("Play#1#Coin", "127.0.0.4",
		testPlayer, playerset.Players{testPlayer}, &deck.Cards{}, &deck.Cards{}, card.Coin)
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Play phase: %v", err)
	}
}

func TestPlayDoNoErrInRoundEnd(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6}
	testObject := NewPlay("Play#1#Coin", "127.0.0.4", testPlayer,
		playerset.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer},
		&testPlayedCards, &deck.Cards{}, card.Coin)
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Play phase: %v", err)
	}
}

func TestPlayWithSideDoNoErr(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testObject := NewPlay("Play#1#Coin", "127.0.0.4",
		testPlayer, playerset.Players{testPlayer}, &deck.Cards{}, nil, card.Coin)
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Play phase: %v", err)
	}
}

func TestPlayWithSideDoNoErrInRoundEnd(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6}
	testObject := NewPlay("Play#1#Coin", "127.0.0.4", testPlayer,
		playerset.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer},
		&testPlayedCards, &deck.Cards{}, card.Coin)
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Play phase: %v", err)
	}
}

func TestPlayWithSideInGameEndSideDeckIsEmpty(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6}
	testSideDeck := deck.Cards{13, 23, 11, 21, 30}
	testObject := NewPlay("Play#1#Coin", "127.0.0.4", testPlayer,
		playerset.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer},
		&testPlayedCards, &testSideDeck, card.Coin)
	testObject.Do(testPlayer)
	if len(testSideDeck) != 0 {
		t.Log(testSideDeck)
		t.Fatal("At the end of the game side deck should be empty")
	}
}

func TestPlayWithSideInGameEndRoundWinningPlayerTakesSideDeck(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayedCards := deck.Cards{2, 3, 4, 6}
	testSideDeck := deck.Cards{13, 23, 11, 21, 30}
	testObject := NewPlay("Play#1#Coin", "127.0.0.4", testPlayer,
		playerset.Players{testPlayer, player.New(), player.New(), player.New(), player.New()},
		&testPlayedCards, &testSideDeck, card.Coin)
	testObject.Do(testPlayer)
	if testPlayer.Count(func(card.ID) uint8 { return 1 }) != 10 {
		t.Fatal("At the end of the game side deck should be empty")
	}
}
