package game

import (
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func postRequest(g roundInformer, rq requestInformer) {
	current := g.Phase()
	switch current {
	case phase.PlayingCards:
		roundHasEnded := len(*g.PlayedCards()) == 5
		if !roundHasEnded {
			break
		}
		playerIndex, _ := g.Players().Find(func(pl *player.Player) bool { return pl == g.CurrentPlayer() })
		winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
		winnerIndex := (playerIndex + int(winningCardIndex) + 1) % 5
		g.Players()[winnerIndex].Collect(g.PlayedCards())
		if !(team.Count(g.Players(), player.IsHandEmpty) == 5 && g.IsSideUsed()) {
			break
		}
		side := g.SideDeck()
		g.Players()[winnerIndex].Collect(side)
		side.Clear()
	}
}
