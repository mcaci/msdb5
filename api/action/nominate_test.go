package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestNominatePhase(t *testing.T) {
	if testObject := NewCompanion("", "", nil, nil, nil); testObject.Phase() != 2 {
		t.Fatalf("Unexpected phase")
	}
}

func TestNominateFindsPlayerInTurn(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "127.0.0.3")
	if testObject := NewCompanion("", "127.0.0.3", testPlayer, nil, nil); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestNominateDoesNotFindPlayerNotInTurn(t *testing.T) {
	testPlayer := player.New()
	if testObject := NewCompanion("", "", nil, nil, nil); testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestNominateNextPlayerOf2is2(t *testing.T) {
	if testObject := NewCompanion("", "", nil, nil, nil); testObject.NextPlayer(2) != 2 {
		t.Fatalf("Next player should be 2")
	}
}

func TestNominateNextPlayerOf4is4(t *testing.T) {
	if testObject := NewCompanion("", "", nil, nil, nil); testObject.NextPlayer(4) != 4 {
		t.Fatalf("Next player should be 1")
	}
}

func TestNominateNextPhaseWithPlayersWithEmptyNameIsTrue(t *testing.T) {
	testPlayers := playerset.Players{player.New()}
	if testObject := NewCompanion("", "", nil, nil, nil); !testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should always be true")
	}
}

func TestNominateNextPhaseWithPlayersWithNonEmptyNameIsFalse(t *testing.T) {
	if testObject := NewCompanion("", "", nil, nil, nil); !testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should always be true")
	}
}
