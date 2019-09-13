package team

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
)

type mockCallers struct {
	call *player.Player
}

func (m mockCallers) Caller() *player.Player  { return m.call }
func (mockCallers) Companion() *player.Player { return player.New() }

func TestTeamCallers(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if !IsInCallers(mockCallers{fakePlayer}, fakePlayer) {
		t.Fatal("Player should be in Callers")
	}
}

func TestTeamOthers(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if IsInCallers(mockCallers{}, fakePlayer) {
		t.Fatal("Player should be in Others")
	}
}
