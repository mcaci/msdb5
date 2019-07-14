package game

import (
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
)

func nextPlayer(g roundInformer, rq requestInformer) {
	current := g.Phase()
	actingPlayerIndex := senderIndex(g, rq)
	var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
	playerIndex := uint8(actingPlayerIndex)
	nextPlayer := playersRoundRobin(playerIndex)
	switch current {
	case phase.ChoosingCompanion, phase.ExchangingCards:
		nextPlayer = playerIndex
	case phase.InsideAuction:
		if *g.AuctionScore() >= 120 {
			for i := range g.Players() {
				if i == actingPlayerIndex {
					continue
				}
				g.Players()[i].Fold()
			}
		}
		for player.Folded(g.Players()[nextPlayer]) {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
	case phase.PlayingCards:
		roundHasEnded := len(*g.PlayedCards()) == 5
		if roundHasEnded {
			winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
			nextPlayer = playersRoundRobin(playerIndex + winningCardIndex)
		}
	default:
	}
	track.Player(g.LastPlaying(), g.Players()[nextPlayer])
}
