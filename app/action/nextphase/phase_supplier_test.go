package nextphase

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/game"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func testObject(phase game.Phase, players team.Players, sideDeck bool, request string) action.NextPhaseChanger {
	return NewChanger(phase, players, sideDeck, request)
}

func TestJoiningStaysInJoiningIfConditionNotMet(t *testing.T) {
	testPlayers := team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(game.Joining, testPlayers, true, "").NextPhase()
	testExpected := game.Joining
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestJoiningGoesToAuctionIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "any")
	testPlayers := team.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(game.Joining, testPlayers, true, "").NextPhase()
	testExpected := game.InsideAuction
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestInsideAuctionStaysIfConditionNotMet(t *testing.T) {
	testPlayer := player.New()
	testPlayers := team.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(game.InsideAuction, testPlayers, true, "").NextPhase()
	testExpected := game.InsideAuction
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestAuctionGoesToCompanionIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Fold()
	testPlayers := team.Players{player.New(), testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(game.InsideAuction, testPlayers, false, "").NextPhase()
	testExpected := game.ChosingCompanion
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestAuctionGoesToExchangeIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Fold()
	testPlayers := team.Players{player.New(), testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(game.InsideAuction, testPlayers, true, "").NextPhase()
	testExpected := game.ExchangingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestExchangingCardsStaysIfConditionNotMet(t *testing.T) {
	testPhase := testObject(game.ExchangingCards, nil, true, "").NextPhase()
	testExpected := game.ExchangingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestExchangingCardsGoesToCompanionIfConditionMet(t *testing.T) {
	testPhase := testObject(game.ExchangingCards, nil, true, "Companion#0").NextPhase()
	testExpected := game.ChosingCompanion
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestCompanionImmediatelyGoesToPlay(t *testing.T) {
	testPhase := testObject(game.ChosingCompanion, nil, true, "").NextPhase()
	testExpected := game.PlayingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestPlayCardStaysIfConditionNotMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayers := team.Players{testPlayer, player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(game.PlayingCards, testPlayers, true, "").NextPhase()
	testExpected := game.PlayingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestPlayCardGoesToEndIfConditionMet(t *testing.T) {
	testPlayers := team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(game.PlayingCards, testPlayers, true, "").NextPhase()
	testExpected := game.End
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}
