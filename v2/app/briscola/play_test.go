package briscola_test

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/app/misc"
	briscolad "github.com/mcaci/msdb5/v2/dom/briscola"
)

type inTest struct {
	name    string
	cardIdx uint8
}

func (i inTest) Name() string { return i.name }
func (i inTest) Idx() uint8   { return i.cardIdx }

func TestEmptyPlayWithError(t *testing.T) {
	in := inTest{"playername", 0}
	gamestate := briscola.Game{}
	err := briscola.Play(&gamestate, in)
	if err == nil {
		t.Error("Expecting an error but all went fine")
	}
}

func TestPlayActionOk(t *testing.T) {
	in := inTest{"playername", 0}
	pl := misc.New(&misc.Options{Name: "playername"})
	pl.Hand().Add(*card.MustID(1))
	gamestate := briscola.Game{PlayerList: misc.Players{pl}, BoardSet: &briscolad.PlayedCards{Cards: &set.Cards{}}}
	err := briscola.Play(&gamestate, in)
	if err != nil {
		t.Error(err)
	}
}
