package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

// PlayData func
func PlayData(g *game.Game, request, origin string) Data {
	phase := game.PlayBriscola
	find := func(p *player.Player) bool { return p.IsExpectedPlayer(g.PlayerInTurn(), origin) }
	do := func(p *player.Player) (err error) {
		data := strings.Split(request, "#")
		number := data[1]
		seed := data[2]
		c, err := card.Create(number, seed)
		p.Play(c)
		g.Board().PlayedCards().Add(c)
		roundMayEnd := len(*g.Board().PlayedCards()) >= 4
		if roundMayEnd && g.CurrentPhase() == game.PlayBriscola {
			roundWinnerIndex := roundWinner(g)
			g.Players()[roundWinnerIndex].Collect(g.Board().PlayedCards())
		}
		return
	}
	nextPlayerOperator := func(playerInTurn uint8) (next uint8) {
		next = nextPlayerInTurn(playerInTurn)
		roundMayEnd := len(*g.Board().PlayedCards()) >= 4
		if roundMayEnd && g.CurrentPhase() == game.PlayBriscola {
			next = roundWinner(g)
			g.Board().PlayedCards().Clear()
		}
		return
	}
	nextPhasePredicate := endGameCondition
	playerPredicate := func(p *player.Player) bool { return p.IsHandEmpty() }
	return Data{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func roundWinner(g *game.Game) uint8 {
	return (g.PlayerInTurnIndex() + briscola.IndexOfWinningCard(*g.Board().PlayedCards(), g.BriscolaSeed()) + 1) % 5
}

func endGameCondition(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.All(searchCriteria)
}
