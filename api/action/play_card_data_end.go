package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// PlayEndRoundData func
func PlayEndRoundData(g *game.Game, request, origin string) Data {
	phase := game.PlayBriscola
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g.PlayerInTurn(), origin) }
	do := func(p *player.Player) (err error) {
		data := strings.Split(request, "#")
		number := data[1]
		seed := data[2]
		c, err := card.Create(number, seed)
		p.Play(c)
		g.Board().PlayedCards().Add(c)
		roundWinnerIndex := roundWinner(g)
		g.Players()[roundWinnerIndex].Collect(g.Board().PlayedCards())
		return
	}
	nextPlayerOperator := func(uint8) uint8 {
		roundWinnerIndex := roundWinner(g)
		g.Board().PlayedCards().Clear()
		return roundWinnerIndex
	}
	nextPhasePredicate := endGameCondition
	playerPredicate := func(p *player.Player) bool { return p.IsHandEmpty() }
	return Data{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func roundWinner(g *game.Game) uint8 {
	return (g.PlayerInTurnIndex() + briscola.IndexOfWinningCard(*g.Board().PlayedCards(), g.BriscolaSeed()) + 1) % 5
}
