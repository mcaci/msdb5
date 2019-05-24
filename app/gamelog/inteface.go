package gamelog

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type informer interface {
	AuctionScore() auction.Score
	Companion() *player.Player
	CurrentPlayer() *player.Player
	IsSideUsed() bool
	LastCardPlayed() card.ID
	LastPlayer() *player.Player
	Phase() phase.ID
	SideDeck() deck.Cards
}
