package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestPlayNextPlayerOf2is3WithRoundNotEnded(t *testing.T) {
	testObject := NewPlay("", "", nil, &deck.Cards{}, nil, card.Coin)
	if nextPlayer := testObject.NextPlayer(2); nextPlayer != 3 {
		t.Fatalf("Next player should be 3, but is %d", nextPlayer)
	}
}

func TestPlayNextPlayerOfAnyIs3WithRoundEnded(t *testing.T) {
	testPlayedCards := deck.Cards{2, 3, 4, 1, 6}
	testObject := NewPlay("", "", nil, &testPlayedCards, nil, card.Coin)
	if nextPlayer := testObject.NextPlayer(0); nextPlayer != 4 {
		t.Fatalf("Next player should be 4, but is %d", nextPlayer)
	}
}

func TestPlayNextPhaseIsFalse(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewPlay("", "", nil, nil, nil, card.Coin); game.End != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should step to end phase")
	}
}

func TestPlayNextPhaseIsTrue(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayers := playerset.Players{player.New(), testPlayer, player.New(), player.New(), player.New()}
	if testObject := NewPlay("", "", nil, nil, nil, card.Coin); game.End == testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in play phase")
	}
}

func TestPlayNextPhaseWithPlayersWithNonFoldedNameIsTrue(t *testing.T) {
	if testObject := NewPlay("", "", nil, nil, nil, card.Coin); !testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should be true with empty handed player")
	}
}

func TestPlayNextPhaseWithFoldedPlayerIsFalse(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	if testObject := NewPlay("", "", nil, nil, nil, card.Coin); testObject.NextPhasePlayerInfo(testPlayer) {
		t.Fatalf("Should be false with non empty handed player")
	}
}
