package briscola5

import (
	"testing"

	"github.com/mcaci/msdb5/v2/dom/player"
)

type callerstest struct{ caller, companion *player.Player }

func NewCallersTeam(clr, cmp *player.Player) *callerstest {
	return &callerstest{caller: clr, companion: cmp}
}
func (c callerstest) Caller() *player.Player    { return c.caller }
func (c callerstest) Companion() *player.Player { return c.companion }

func TestTeamCallers(t *testing.T) {
	fakePlayer := NewPlayer()
	if !IsInCallers(NewCallersTeam(&fakePlayer.Player, &NewPlayer().Player))(&fakePlayer.Player) {
		t.Fatal("Player should be in Callers")
	}
}

func TestTeamOthers(t *testing.T) {
	fakePlayer := NewPlayer()
	if IsInCallers(NewCallersTeam(&NewPlayer().Player, &NewPlayer().Player))(&fakePlayer.Player) {
		t.Fatal("Player should be in Others")
	}
}
