package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) playEndRoundData(request, origin string) phaseData {
	phase := playBriscola
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g, origin) }
	do := func(p *player.Player) (err error) {
		data := strings.Split(request, "#")
		number := data[1]
		seed := data[2]
		c, err := card.Create(number, seed)
		p.Play(c)
		g.info.PlayedCards().Add(c)
		roundWinnerIndex := roundWinner(g)
		g.players[roundWinnerIndex].Collect(g.info.PlayedCards())
		return
	}
	nextPlayerOperator := func(uint8) uint8 {
		roundWinnerIndex := roundWinner(g)
		g.info.PlayedCards().Clear()
		return roundWinnerIndex
	}
	nextPhasePredicate := g.endGameCondition
	playerPredicate := func(p *player.Player) bool { return p.IsHandEmpty() }
	return phaseData{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func roundWinner(g *Game) uint8 {
	return (g.playerInTurn + briscola.IndexOfWinningCard(*g.info.PlayedCards(), g.companion.Card().Seed()) + 1) % 5
}
