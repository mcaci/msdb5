package game

import (
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
)

func nextPlayer(g roundInformer, rq interface{ From() string }) uint8 {
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
		if !g.IsRoundOngoing() {
			winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
			nextPlayer = playersRoundRobin(playerIndex + winningCardIndex)
		}
	}
	return nextPlayer
}
