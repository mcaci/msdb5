package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestNominateWithSideDoNoErr(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testObject := NewCompanionWithSide("Companion#1#Coin", "127.0.0.2",
		testPlayer, playerset.Players{testPlayer}, func(card.ID, *player.Player) {
			return
		})
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Nominate phase: %v", err)
	}
}

func TestNominateWithSideDoErr_PlayerWithCardNotFound(t *testing.T) {
	testPlayer := player.New()
	testObject := NewCompanionWithSide("Companion#1#Coin", "127.0.0.2",
		testPlayer, playerset.Players{}, func(card.ID, *player.Player) {
			return
		})
	err := testObject.Do(testPlayer)
	if err == nil {
		t.Fatalf("Unexpected error from Nominate phase: %v", err)
	}
}

func TestNominateWithSideDoErr_CardNotExistent(t *testing.T) {
	testPlayer := player.New()
	testObject := NewCompanionWithSide("Companion#1#Coins", "127.0.0.2",
		testPlayer, playerset.Players{testPlayer}, func(card.ID, *player.Player) {
			return
		})
	err := testObject.Do(testPlayer)
	if err == nil {
		t.Fatalf("Unexpected error from Nominate phase: %v", err)
	}
}
