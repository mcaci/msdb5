package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestJoinPhase(t *testing.T) {
	if testObject := NewJoin("", ""); testObject.Phase() != 0 {
		t.Fatalf("Unexpected phase")
	}
}

func TestJoinFindsPlayerWithEmptyName(t *testing.T) {
	testPlayer := player.New()
	if testObject := NewJoin("", ""); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestJoinDoesNotFindPlayerWithEmptyName(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "")
	if testObject := NewJoin("", ""); testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestJoinNextPlayerOf0is1(t *testing.T) {
	if testObject := NewJoin("", ""); testObject.NextPlayer(0) != 1 {
		t.Fatalf("Next player should be 1")
	}
}

func TestJoinNextPlayerOf4is0(t *testing.T) {
	if testObject := NewJoin("", ""); testObject.NextPlayer(4) != 0 {
		t.Fatalf("Next player should be 1")
	}
}

func TestJoinNextPhaseWithPlayersWithEmptyNameIsTrue(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "")
	testPlayers := playerset.Players{testPlayer}
	if testObject := NewJoin("", ""); 1 != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Unexpected play next phase")
	}
}

func TestJoinNextPhaseWithPlayersWithNonEmptyNameIsFalse(t *testing.T) {
	testPlayer := player.New()
	testPlayers := playerset.Players{testPlayer}
	if testObject := NewJoin("", ""); 1 == testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Unexpected play next phase")
	}
}
