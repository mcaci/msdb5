package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestNominateWithSidePhase(t *testing.T) {
	if testObject := NewCompanionWithSide("", "", nil, nil, nil); testObject.Phase() != 2 {
		t.Fatalf("Unexpected phase")
	}
}

func TestNominateWithSideFindsPlayerInTurn(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "127.0.0.3")
	if testObject := NewCompanionWithSide("", "127.0.0.3", testPlayer, nil, nil); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestNominateWithSideDoesNotFindPlayerNotInTurn(t *testing.T) {
	testPlayer := player.New()
	if testObject := NewCompanionWithSide("", "", nil, nil, nil); testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestNominateWithSideNextPlayerOf2is2(t *testing.T) {
	if testObject := NewCompanionWithSide("", "", nil, nil, nil); testObject.NextPlayer(2) != 2 {
		t.Fatalf("Next player should be 2")
	}
}

func TestNominateWithSideNextPlayerOf4is4(t *testing.T) {
	if testObject := NewCompanionWithSide("", "", nil, nil, nil); testObject.NextPlayer(4) != 4 {
		t.Fatalf("Next player should be 1")
	}
}

func TestNominateWithSideNextPhaseWithPlayersWithEmptyNameIsTrue(t *testing.T) {
	testPlayers := playerset.Players{player.New()}
	if testObject := NewCompanionWithSide("", "", nil, nil, nil); 3 != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should always be true")
	}
}

func TestNominateWithSideNextPhaseWithPlayersWithNonEmptyNameIsFalse(t *testing.T) {
	if testObject := NewCompanionWithSide("", "", nil, nil, nil); !testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should always be true")
	}
}
