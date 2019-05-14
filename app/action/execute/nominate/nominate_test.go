package nominate

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func TestNominateDoNoErr(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testObject := NewCompanion("Companion#1#Coin", "127.0.0.2",
		team.Players{testPlayer}, func(card.ID, *player.Player) {
			return
		})
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Nominate phase: %v", err)
	}
}

func TestNominateDoErr_PlayerWithCardNotFound(t *testing.T) {
	testPlayer := player.New()
	testObject := NewCompanion("Companion#1#Coin", "127.0.0.2",
		team.Players{}, func(card.ID, *player.Player) {
			return
		})
	err := testObject.Do(testPlayer)
	if err == nil {
		t.Fatalf("Unexpected error from Nominate phase: %v", err)
	}
}

func TestNominateDoErr_CardNotExistent(t *testing.T) {
	testPlayer := player.New()
	testObject := NewCompanion("Companion#1#Coins", "127.0.0.2",
		team.Players{testPlayer}, func(card.ID, *player.Player) {
			return
		})
	err := testObject.Do(testPlayer)
	if err == nil {
		t.Fatalf("Unexpected error from Nominate phase: %v", err)
	}
}
