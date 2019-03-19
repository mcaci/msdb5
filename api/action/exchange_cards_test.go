package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestExchangeCardsPhase(t *testing.T) {
	if testObject := NewExchangeCards("", "", nil); testObject.Phase() != game.ExchangingCards {
		t.Fatalf("Unexpected phase")
	}
}

func TestExchangeCardsFindsPlayerInTurn(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "127.0.0.3")
	if testObject := NewExchangeCards("", "127.0.0.3", testPlayer); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestExchangeCardsNextPlayerOf2is2(t *testing.T) {
	if testObject := NewExchangeCards("", "", nil); testObject.NextPlayer(2) != 2 {
		t.Fatalf("Next player should be 2")
	}
}

func TestExchangeCardsNextPlayerOf4is4(t *testing.T) {
	if testObject := NewExchangeCards("", "", nil); testObject.NextPlayer(4) != 4 {
		t.Fatalf("Next player should be 1")
	}
}

func TestExchangeCardsNextPhaseWhenInputIs0(t *testing.T) {
	testPlayers := playerset.Players{player.New()}
	if testObject := NewExchangeCards("#0", "", nil); game.ChosingCompanion != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should change phase when 0 is in the request")
	}
}

func TestExchangeCardsNextPhaseWhenInputIsNot0(t *testing.T) {
	testPlayers := playerset.Players{player.New()}
	if testObject := NewExchangeCards("#1", "", nil); game.ExchangingCards != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should not change phase when 1 is in the request")
	}
}

func TestExchangeCardsNextPhaseWithPlayersWithNonEmptyNameIsFalse(t *testing.T) {
	if testObject := NewExchangeCards("", "", nil); !testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should always be true")
	}
}
