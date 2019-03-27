package nextphase

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/game"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/playerset"
)

func testObject(phase game.Phase, players playerset.Players, sideDeck bool, request string) action.NextPhaseChanger {
	return NewChanger(phase, players, sideDeck, request)
}

func TestJoiningStaysInJoiningIfConditionNotMet(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(game.Joining, testPlayers, true, "").NextPhase()
	testExpected := game.Joining
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestJoiningGoesToAuctionIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "any")
	testPlayers := playerset.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(game.Joining, testPlayers, true, "").NextPhase()
	testExpected := game.InsideAuction
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestInsideAuctionStaysIfConditionNotMet(t *testing.T) {
	testPlayer := player.New()
	testPlayers := playerset.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(game.InsideAuction, testPlayers, true, "").NextPhase()
	testExpected := game.InsideAuction
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestAuctionGoesToCompanionIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Fold()
	testPlayers := playerset.Players{player.New(), testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(game.InsideAuction, testPlayers, false, "").NextPhase()
	testExpected := game.ChosingCompanion
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestAuctionGoesToExchangeIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Fold()
	testPlayers := playerset.Players{player.New(), testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(game.InsideAuction, testPlayers, true, "").NextPhase()
	testExpected := game.ExchangingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestExchangingCardsStaysIfConditionNotMet(t *testing.T) {
	testPhase := testObject(game.ExchangingCards, nil, true, "").NextPhase()
	testExpected := game.ExchangingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestExchangingCardsGoesToCompanionIfConditionMet(t *testing.T) {
	testPhase := testObject(game.ExchangingCards, nil, true, "Companion#0").NextPhase()
	testExpected := game.ChosingCompanion
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestCompanionImmediatelyGoesToPlay(t *testing.T) {
	testPhase := testObject(game.ChosingCompanion, nil, true, "").NextPhase()
	testExpected := game.PlayingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestPlayCardStaysIfConditionNotMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayers := playerset.Players{testPlayer, player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(game.PlayingCards, testPlayers, true, "").NextPhase()
	testExpected := game.PlayingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}

func TestPlayCardGoesToEndIfConditionMet(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(game.PlayingCards, testPlayers, true, "").NextPhase()
	testExpected := game.End
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d", testExpected)
	}
}
