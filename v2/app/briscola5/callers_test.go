package briscola5

import (
	"testing"

	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

type callerstest struct{ caller, companion misc.Player }

func NewCallersTeam(clr, cmp misc.Player) *callerstest {
	return &callerstest{caller: clr, companion: cmp}
}
func (c callerstest) Caller() misc.Player    { return c.caller }
func (c callerstest) Companion() misc.Player { return c.companion }

func TestTeamCallers(t *testing.T) {
	p := briscola5.NewB5Player("p")
	if !IsInCallers(NewCallersTeam(p, briscola5.NewB5Player("q")))(p) {
		t.Fatal("misc.Player should be in Callers")
	}
}

func TestTeamOthers(t *testing.T) {
	p := briscola5.NewB5Player("p")
	if IsInCallers(NewCallersTeam(p, p))(briscola5.NewB5Player("q")) {
		t.Fatal("misc.Player should be in Others")
	}
}
