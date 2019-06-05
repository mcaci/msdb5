package game

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
)

func nextPlayer(g *Game, request, origin string, current phase.ID) {
	actingPlayerIndex := g.senderIndex(origin)
	var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
	playerIndex := uint8(actingPlayerIndex)
	nextPlayer := playersRoundRobin(playerIndex)
	switch current {
	case phase.ChoosingCompanion, phase.ExchangingCards:
		nextPlayer = playerIndex
	case phase.InsideAuction:
		for g.players[nextPlayer].Folded() {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
		if nextPlayer == playerIndex {
			g.caller = g.players[playerIndex]
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
