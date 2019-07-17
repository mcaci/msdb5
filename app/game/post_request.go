package game

import (
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/deck"
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
		p := g.Players()[winnerIndex]
		move(g.PlayedCards(), p.Pile())
		if !(team.Count(g.Players(), player.IsHandEmpty) == 5 && g.IsSideUsed()) {
			break
		}
		move(g.SideDeck(), p.Pile())
	}
}

func move(from, to *deck.Cards) {
	to.Add(*from...)
	from.Clear()
}
