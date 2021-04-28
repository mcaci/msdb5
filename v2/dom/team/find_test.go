package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/player"
)

func TestUnsuccessfulFind(t *testing.T) {
	if _, err := testPlayers.Index(player.IsCardInHand(*card.MustID(8))); err == nil {
		t.Fatal("Player should not be found")
	}
}

func TestSuccessfulFindNoErr(t *testing.T) {
	if _, err := testPlayers.Index(player.IsCardInHand(*card.MustID(33))); err != nil {
		t.Fatal("Player not found with criteria player.IsCardInHand(33)")
	}
}

func TestSuccessfulFindWithNone(t *testing.T) {
	if testPlayers.None(player.IsCardInHand(*card.MustID(33))) {
		t.Fatal("Player not found with criteria player.IsCardInHand(33)")
	}
}

func TestUnsuccessfulFindWithNone(t *testing.T) {
	if !testPlayers.None(player.IsCardInHand(*card.MustID(24))) {
		t.Fatal("Player should not be found")
	}
}

func TestUnsuccessfulFindWithAll(t *testing.T) {
	if testPlayers.All(player.IsCardInHand(*card.MustID(24))) {
		t.Fatal("Player should not be found")
	}
}

func TestSuccessfulFindWithAll(t *testing.T) {
	if !testPlayers.All(player.IsCardInHand(*card.MustID(34))) {
		t.Fatal("All players have the 4 of Cudgel")
	}
}
