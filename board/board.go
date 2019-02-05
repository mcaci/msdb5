package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	players      player.Players
	pChans       []chan card.ID
	playedCards  card.Cards
	selectedCard card.ID
	auctionScore uint8
}

// New func
func New() *Board {
	b := new(Board)
	b.pChans = make([]chan card.ID, 5)
	for i := range b.pChans {
		b.pChans[i] = make(chan card.ID)
	}
	makePlayers(b)
	playersDrawAllCards(&b.players)
	return b
}

// NewAPI func
func NewAPI() API {
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
