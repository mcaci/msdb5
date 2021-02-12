package collect

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/phase"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type testObj phase.ID

func (t testObj) Phase() phase.ID { return phase.ID(t) }

func TestNoCollectorFunc(t *testing.T) {
	cards := &set.Cards{*card.MustID(1), *card.MustID(2)}
	side := &set.Cards{*card.MustID(4), *card.MustID(5)}
	ph := testObj(phase.ChoosingCompanion)
	all := team.Players{player.New()}
	col := Collector(ph, all, side, cards)
	if len(*col()) != 0 {
		t.Fatal("Unexpected cards returned")
	}
}

func TestCollectorFunc(t *testing.T) {
	cards := &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(16), *card.MustID(32), *card.MustID(11)}
	side := &set.Cards{*card.MustID(4), *card.MustID(5), *card.MustID(12), *card.MustID(25), *card.MustID(23)}
	ph := testObj(phase.End)
	all := team.Players{player.New()}
	col := Collector(ph, all, side, cards)
	if len(*col()) == 0 {
		t.Fatal("Unexpected cards not returned")
	}
}
