package team

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
)

func TestTeamCallers(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if !IsInCallers(NewCallers(fakePlayer), fakePlayer) {
		t.Fatal("Player should be in Callers")
	}
}

func TestTeamOthers(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if IsInCallers(NewEmptyCallers(), fakePlayer) {
		t.Fatal("Player should be in Others")
	}
}
