package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"

	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestJoinNextPlayerOf0is1(t *testing.T) {
	if testObject := NewPlayerSelector(); testObject.NextPlayer(0) != 1 {
		t.Fatalf("Next player should be 1")
	}
}

func TestJoinNextPlayerOf4is0(t *testing.T) {
	if testObject := NewPlayerSelector(); testObject.NextPlayer(4) != 0 {
		t.Fatalf("Next player should be 1")
	}
}

func TestJoinNextPhaseWithPlayersWithEmptyNameIsTrue(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "")
	testPlayers := playerset.Players{testPlayer}
	if testObject := NewPhaseChanger(testPlayers); game.InsideAuction != testObject.NextPhase() {
		t.Fatalf("Unexpected play next phase")
	}
}

func TestJoinNextPhaseWithPlayersWithNonEmptyNameIsFalse(t *testing.T) {
	testPlayer := player.New()
	testPlayers := playerset.Players{testPlayer}
	if testObject := NewPhaseChanger(testPlayers); game.InsideAuction == testObject.NextPhase() {
		t.Fatalf("Unexpected play next phase")
	}
}