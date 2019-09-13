package msg

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

// TranslateTeam func
func TranslateTeam(p *player.Player, g team.Callers, printer *message.Printer) string {
	t := printer.Sprintf("Callers")
	if !team.IsInCallers(g, p) {
		t = printer.Sprintf("Others")
	}
	return printer.Sprintf("The end - %s team has all briscola cards\n", t)
}

type selfInformer interface {
	Phase() phase.ID
	SideDeck() *set.Cards
}

// TranslatePlayer func
func TranslatePlayer(pl *player.Player, g interface{ Briscola() card.Item }, printer *message.Printer) string {
	var seed *card.Seed
	if g.Briscola().Number() > 0 {
		s := g.Briscola().Seed()
		seed = &s
	}
	return printer.Sprintf("Player: (Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)\n",
		pl.Name(), TranslateHand(*pl.Hand(), seed, printer), TranslateCards(*pl.Pile(), printer), player.Folded(pl))
}
