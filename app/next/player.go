package next

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type nextPlayerInformer interface {
	Players() team.Players
	PlayedCards() *set.Cards
	Phase() phase.ID
	Briscola() card.Item
	IsSideUsed() bool
	IsRoundOngoing() bool
	FromInput() string
}

// Player func
func Player(g nextPlayerInformer) *player.Player {
	numberOfPlayers := uint8(len(g.Players()))
	playersRoundRobin := func(playerIndex uint8) uint8 { return (playerIndex + 1) % numberOfPlayers }
	index, _ := g.Players().Find(player.MatchingHost(g.FromInput()))
	playerIndex := uint8(index)
	nextPlayer := playersRoundRobin(playerIndex)
	switch g.Phase() {
	case phase.InsideAuction:
		for player.Folded(g.Players()[nextPlayer]) {
			nextPlayer = playersRoundRobin(nextPlayer)
		}
	case phase.ChoosingCompanion, phase.ExchangingCards:
		nextPlayer = playerIndex
	case phase.PlayingCards:
		if g.IsRoundOngoing() {
			break
		}
		winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
		nextPlayer = playersRoundRobin(playerIndex + winningCardIndex)
	}
	return g.Players()[nextPlayer]
}
