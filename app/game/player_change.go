package game

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
)

func nextPlayer(g *Game, current phase.ID, playerInTurn uint8) uint8 {
	var playersRoundRobin = func(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }
	nextPlayer := playersRoundRobin(playerInTurn)
	switch current {
	case phase.ChosingCompanion, phase.ExchangingCards:
		nextPlayer = playerInTurn
	case phase.InsideAuction:
		for g.players[nextPlayer].Folded() {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
	case phase.PlayingCards:
		roundHasEnded := len(g.playedCards) == 5
		if roundHasEnded {
			winningCardIndex := briscola.IndexOfWinningCard(g.playedCards, g.BriscolaSeed())
			nextPlayer = playersRoundRobin(playerInTurn + winningCardIndex)
		}
	default:
	}
	return nextPlayer
}
