package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestPlayWithSidePhase(t *testing.T) {
	if testObject := NewPlayWithSide("", "", nil, nil, nil, nil, card.Coin); testObject.Phase() != game.PlayingCards {
		t.Fatalf("Unexpected phase")
	}
}

func TestPlayWithSideFindsPlayerInTurn(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "127.0.0.4")
	if testObject := NewPlayWithSide("", "127.0.0.4", testPlayer, nil, nil, nil, card.Coin); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestPlayWithSideDoesNotFindPlayerNotInTurn(t *testing.T) {
	if testObject := NewPlayWithSide("", "", nil, nil, nil, nil, card.Coin); testObject.Find(player.New()) {
		t.Fatalf("Unexpected player")
	}
}

func TestPlayWithSideNextPlayerOf2is3WithRoundNotEnded(t *testing.T) {
	testObject := NewPlayWithSide("", "", nil, nil, &deck.Cards{}, nil, card.Coin)
	if nextPlayer := testObject.NextPlayer(2); nextPlayer != 3 {
		t.Fatalf("Next player should be 3, but is %d", nextPlayer)
	}
}

func TestPlayWithSideNextPlayerOfAnyIs3WithRoundEnded(t *testing.T) {
	testPlayedCards := deck.Cards{2, 3, 4, 1, 6}
	testObject := NewPlayWithSide("", "", nil, nil, &testPlayedCards, nil, card.Coin)
	if nextPlayer := testObject.NextPlayer(0); nextPlayer != 4 {
		t.Fatalf("Next player should be 4, but is %d", nextPlayer)
	}
}

func TestPlayWithSideNextPhaseIsFalse(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewPlayWithSide("", "", nil, nil, nil, nil, card.Coin); game.End != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should step to end phase")
	}
}

func TestPlayWithSideNextPhaseIsTrue(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayers := playerset.Players{player.New(), testPlayer, player.New(), player.New(), player.New()}
	if testObject := NewPlayWithSide("", "", nil, nil, nil, nil, card.Coin); game.End == testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in play phase")
	}
}

func TestPlayWithSideNextPhaseWithPlayersWithNonFoldedNameIsTrue(t *testing.T) {
	if testObject := NewPlayWithSide("", "", nil, nil, nil, nil, card.Coin); !testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should be true with empty handed player")
	}
}

func TestPlayWithSideNextPhaseWithFoldedPlayerIsFalse(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	if testObject := NewPlayWithSide("", "", nil, nil, nil, nil, card.Coin); testObject.NextPhasePlayerInfo(testPlayer) {
		t.Fatalf("Should be false with non empty handed player")
	}
}
