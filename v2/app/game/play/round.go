package play

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

type RoundOptions struct {
	PlHand       set.Cards
	Idx          uint8
	PlayedCards  set.Cards
	NPlayers     uint8
	BriscolaCard interface{ Seed() card.Seed }
}

type RoundInfo struct {
	OnBoard set.Cards
	NextPl  uint8
	NextRnd bool
}

func Round(g *RoundOptions) *RoundInfo {
	defaultInfo := &RoundInfo{
		OnBoard: g.PlayedCards,
		NextPl:  roundRobin(g.Idx, 1, g.NPlayers),
	}
	if len(g.PlHand) <= 0 {
		return defaultInfo
	}
	err := set.MoveOne(&g.PlHand[g.Idx], &g.PlHand, &g.PlayedCards)
	if err != nil {
		return defaultInfo
	}
	if !isRoundOngoing(g.PlayedCards) {
		// end current round
		winningCardIndex := indexOfWinningCard(g.PlayedCards, g.BriscolaCard.Seed())
		return &RoundInfo{
			OnBoard: g.PlayedCards,
			NextPl:  roundRobin(g.Idx, winningCardIndex, g.NPlayers),
			NextRnd: true,
		}
	}
	return defaultInfo
}

func roundRobin(idx, off, size uint8) uint8 {
	return (idx + off) % size
}
