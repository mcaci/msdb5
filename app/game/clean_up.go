package game

import (
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func (g *Game) cleanUp(winnerIndex uint8) {
	current := g.Phase()
	if current != phase.PlayingCards || g.IsRoundOngoing() {
		return
	}
	pile := g.Players()[winnerIndex].Pile()
	move(g.PlayedCards(), pile)
	if !(team.Count(g.Players(), player.IsHandEmpty) == 5 && g.IsSideUsed()) {
		return
	}
	move(g.SideDeck(), pile)
}

func move(from, to *deck.Cards) {
	to.Add(*from...)
	from.Clear()
}
