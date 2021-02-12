package team

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/player"
)

// CallerTeam struct
type CallerTeam struct {
	call *player.Player
}

func NewEmptyCallers() Callers               { return CallerTeam{player.New()} }
func NewCallers(call *player.Player) Callers { return CallerTeam{call} }
func (t CallerTeam) Caller() *player.Player  { return t.call }
func (CallerTeam) Companion() *player.Player { return player.New() }

func TestTeamCallers(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if !IsInCallers(NewCallers(fakePlayer))(fakePlayer) {
		t.Fatal("Player should be in Callers")
	}
}

func TestTeamOthers(t *testing.T) {
	fakePlayer := player.New()
	fakePlayer.Pile().Add(*set.NewMust(1)...)
	if IsInCallers(NewEmptyCallers())(fakePlayer) {
		t.Fatal("Player should be in Others")
	}
}
