package action

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type nextPlayerInformer interface {
	Players() team.Players
	PlayedCards() *set.Cards
	Phase() phase.ID
	Briscola() card.Item
	IsRoundOngoing() bool
	FromInput() string
}

func roundRobin(idx, off, size uint8) uint8 {
	return (idx + off) % size
}

// Player func
func Player(g nextPlayerInformer) *player.Player {
	numberOfPlayers := uint8(len(g.Players()))
	index, _ := g.Players().Index(player.MatchingHost(g.FromInput()))
	nextPlayer := roundRobin(index, 1, numberOfPlayers)
	switch g.Phase() {
	case phase.InsideAuction:
		for player.Folded(g.Players()[nextPlayer]) {
			nextPlayer = roundRobin(nextPlayer, 1, numberOfPlayers)
		}
	case phase.ChoosingCompanion, phase.ExchangingCards:
		nextPlayer = index
	case phase.PlayingCards:
		if g.IsRoundOngoing() {
			break
		}
		winningCardIndex := indexOfWinningCard(*g.PlayedCards(), g.Briscola())
		nextPlayer = roundRobin(nextPlayer, winningCardIndex, numberOfPlayers)
	case phase.End:
		if g.IsRoundOngoing() {
			break
		}
		if !player.IsHandEmpty(g.Players()[nextPlayer]) {
			highbriscolaCard := serie(g.Briscola().Seed())
			for _, card := range highbriscolaCard {
				i, err := g.Players().Index(player.IsCardInHand(card))
				if err != nil { // no one has card
					continue
				}
				nextPlayer = i
			}
		}
		winningCardIndex := indexOfWinningCard(*g.PlayedCards(), g.Briscola())
		nextPlayer = roundRobin(nextPlayer, winningCardIndex, numberOfPlayers)
	}
	return g.Players()[nextPlayer]
}

func indexOfWinningCard(cardsOnTheTable set.Cards, b card.Item) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if briscola.IsOtherWinning(base, other, b) {
			base = other
			max = i
		}
	}
	return uint8(max)
}
