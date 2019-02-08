package board

import (
	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	players      player.Players
	playedCards  card.Cards
	selectedCard card.ID
	auctionScore uint8
}

// New func
func New() *Board {
	b := new(Board)
	makePlayers(b)
	playersDrawAllCards(&b.players)
	return b
}

// NewAction func
func NewAction() api.Action {
	return New()
}

func makePlayers(b *Board) {
	for i := 0; i < 5; i++ {
		b.players.Add(*player.New())
	}
}

func playersDrawAllCards(players *player.Players) {
	deck := card.Deck()
	for i := 0; i < card.DeckSize; i++ {
		(*players)[i%5].Draw(&deck)
	}
}
