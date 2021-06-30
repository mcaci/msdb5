package player

import (
	"testing"
)

type callerstest struct{ caller, companion Player }

func NewCallersTeam(clr, cmp Player) *callerstest {
	return &callerstest{caller: clr, companion: cmp}
}
func (c callerstest) Caller() Player    { return c.caller }
func (c callerstest) Companion() Player { return c.companion }

func TestTeamCallers(t *testing.T) {
	p := testP("p")
	if !IsInCallers(NewCallersTeam(p, testP("q")))(p) {
		t.Fatal("Player should be in Callers")
	}
}

func TestTeamOthers(t *testing.T) {
	p := testP("p")
	if IsInCallers(NewCallersTeam(p, p))(testP("q")) {
		t.Fatal("Player should be in Others")
	}
}
