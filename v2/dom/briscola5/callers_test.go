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

func fakePlayer() *player.Player {
	return player.NewWithOpts(&player.Options{For2P: true}).(*player.Player)
}

func TestTeamCallers(t *testing.T) {
	p := fakePlayer()
	if !IsInCallers(NewCallersTeam(p, fakePlayer()))(p) {
		t.Fatal("Player should be in Callers")
	}
}

func TestTeamOthers(t *testing.T) {
	p := fakePlayer()
	if IsInCallers(NewCallersTeam(fakePlayer(), fakePlayer()))(p) {
		t.Fatal("Player should be in Others")
	}
}
