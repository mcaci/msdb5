package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"
)

func TestNominateNextPlayerOf2is2(t *testing.T) {
	if testObject := NewCompanion("", "", nil, nil); testObject.NextPlayer(2) != 2 {
		t.Fatalf("Next player should be 2")
	}
}

func TestNominateNextPlayerOf4is4(t *testing.T) {
	if testObject := NewCompanion("", "", nil, nil); testObject.NextPlayer(4) != 4 {
		t.Fatalf("Next player should be 1")
	}
}

func TestNominateNextPhaseWithPlayersWithEmptyNameIsTrue(t *testing.T) {
	if testObject := NewCompanion("", "", nil, nil); game.PlayingCards != testObject.NextPhase() {
		t.Fatalf("Should always be true")
	}
}