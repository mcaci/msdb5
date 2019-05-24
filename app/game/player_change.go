package game

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
)

func nextPlayer(g *Game, current phase.ID, index int) {
	var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
	playerIndex := uint8(index)
	nextPlayer := playersRoundRobin(playerIndex)
	switch current {
	case phase.ChosingCompanion, phase.ExchangingCards:
		nextPlayer = playerIndex
	case phase.InsideAuction:
		for g.players[nextPlayer].Folded() {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
		if nextPlayer == playerIndex {
			g.caller = g.playersRef()[playerIndex]
		}
	case phase.PlayingCards:
		roundHasEnded := len(g.playedCards) == 5
		if roundHasEnded {
			winningCardIndex := briscola.IndexOfWinningCard(g.playedCards, g.briscola())
			nextPlayer = playersRoundRobin(playerIndex + winningCardIndex)
		}
	default:
	}
	trackActing(&g.lastPlaying, g.players[nextPlayer])
}
