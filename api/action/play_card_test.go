package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestPlayPhase(t *testing.T) {
	if testObject := NewPlay("", "", 0, nil, nil, nil, card.Coin); testObject.Phase() != 3 {
		t.Fatalf("Unexpected phase")
	}
}

func TestPlayFindsPlayerInTurn(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "127.0.0.4")
	if testObject := NewPlay("", "127.0.0.4", 0, testPlayer, nil, nil, card.Coin); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestPlayDoesNotFindPlayerNotInTurn(t *testing.T) {
	if testObject := NewPlay("", "", 0, nil, nil, nil, card.Coin); testObject.Find(player.New()) {
		t.Fatalf("Unexpected player")
	}
}

func TestPlayNextPlayerOf2is3WithRoundNotEnded(t *testing.T) {
	any := uint8(0)
	if testObject := NewPlay("", "", any, nil, nil, board.New(), card.Coin); testObject.NextPlayer(2) != 3 {
		t.Fatalf("Next player should be 3, but is %d", testObject.NextPlayer(any))
	}
}

func TestPlayNextPlayerOfAnyIs3WithRoundEnded(t *testing.T) {
	any := uint8(0)
	testPlayedCards := deck.Cards{2, 3, 4, 1, 6}
	testBoard := board.New()
	testBoard.PlayedCards().Add(testPlayedCards...)
	testObject := NewPlay("", "", any, nil, nil, testBoard, card.Coin)
	if nextPlayer := testObject.NextPlayer(0); nextPlayer != 4 {
		t.Fatalf("Next player should be 4, but is %d", nextPlayer)
	}
}

func TestPlayNextPhaseIsFalse(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewPlay("", "", 0, nil, nil, nil, card.Coin); !testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should step to end phase")
	}
}

func TestPlayNextPhaseIsTrue(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayers := playerset.Players{player.New(), testPlayer, player.New(), player.New(), player.New()}
	if testObject := NewPlay("", "", 0, nil, nil, nil, card.Coin); testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in play phase")
	}
}

func TestPlayNextPhaseWithPlayersWithNonFoldedNameIsTrue(t *testing.T) {
	if testObject := NewPlay("", "", 0, nil, nil, nil, card.Coin); !testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should be true with empty handed player")
	}
}

func TestPlayNextPhaseWithFoldedPlayerIsFalse(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	if testObject := NewPlay("", "", 0, nil, nil, nil, card.Coin); testObject.NextPhasePlayerInfo(testPlayer) {
		t.Fatalf("Should be false with non empty handed player")
	}
}
