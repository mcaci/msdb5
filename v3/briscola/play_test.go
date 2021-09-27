package briscola_test

import (
	"testing"

	"github.com/mcaci/msdb5/v3/briscola"
)

type inTest struct {
	name    string
	cardIdx uint8
}

func (i inTest) Name() string { return i.name }
func (i inTest) Idx() uint8   { return i.cardIdx }

func TestEmptyPlayWithError(t *testing.T) {
	in := inTest{"playername", 0}
	gamestate := briscola.NewGame(briscola.WithDefaultOptions)
	err := briscola.Play(gamestate, in)
	if err == nil {
		t.Error("Expecting an error but all went fine")
	}
}

func TestPlayActionOk(t *testing.T) {
	gamestate := briscola.NewGame(briscola.WithDefaultOptions)
	briscola.Register("playername", gamestate)
	briscola.Register("p2", gamestate)
	err := briscola.Play(gamestate, inTest{"playername", 0})
	if err != nil {
		t.Error(err)
	}
}
