package nextphase

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func testObject(phase phase.ID, players team.Players, sideDeck bool, request string) action.NextPhaseChanger {
	return NewChanger(phase, players, sideDeck, request)
}

func TestJoiningStaysInJoiningIfConditionNotMet(t *testing.T) {
	testPlayers := team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(phase.Joining, testPlayers, true, "").NextPhase()
	testExpected := phase.Joining
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestJoiningGoesToAuctionIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "any")
	testPlayers := team.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(phase.Joining, testPlayers, true, "").NextPhase()
	testExpected := phase.InsideAuction
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestInsideAuctionStaysIfConditionNotMet(t *testing.T) {
	testPlayer := player.New()
	testPlayers := team.Players{testPlayer, testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(phase.InsideAuction, testPlayers, true, "").NextPhase()
	testExpected := phase.InsideAuction
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestAuctionGoesToCompanionIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Fold()
	testPlayers := team.Players{player.New(), testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(phase.InsideAuction, testPlayers, false, "").NextPhase()
	testExpected := phase.ChosingCompanion
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestAuctionGoesToExchangeIfConditionMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Fold()
	testPlayers := team.Players{player.New(), testPlayer, testPlayer, testPlayer, testPlayer}
	testPhase := testObject(phase.InsideAuction, testPlayers, true, "").NextPhase()
	testExpected := phase.ExchangingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestExchangingCardsStaysIfConditionNotMet(t *testing.T) {
	testPhase := testObject(phase.ExchangingCards, nil, true, "").NextPhase()
	testExpected := phase.ExchangingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestExchangingCardsGoesToCompanionIfConditionMet(t *testing.T) {
	testPhase := testObject(phase.ExchangingCards, nil, true, "Companion#0").NextPhase()
	testExpected := phase.ChosingCompanion
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestCompanionImmediatelyGoesToPlay(t *testing.T) {
	testPhase := testObject(phase.ChosingCompanion, nil, true, "").NextPhase()
	testExpected := phase.PlayingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestPlayCardStaysIfConditionNotMet(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testPlayers := team.Players{testPlayer, player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(phase.PlayingCards, testPlayers, true, "").NextPhase()
	testExpected := phase.PlayingCards
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}

func TestPlayCardGoesToEndIfConditionMet(t *testing.T) {
	testPlayers := team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	testPhase := testObject(phase.PlayingCards, testPlayers, true, "").NextPhase()
	testExpected := phase.End
	if testExpected != testPhase {
		t.Fatalf("Should be in phase %d but is in phase %d", testExpected, testPhase)
	}
}
