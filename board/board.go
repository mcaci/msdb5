package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	players      []*player.Player
	playedCards  set.Cards
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

func makePlayers(b *Board) {
	for i := 0; i < 5; i++ {
		b.players = append(b.players, player.New())
	}
}

func playersDrawAllCards(players *[]*player.Player) {
	deck := set.Deck()
	for i := 0; i < set.DeckSize; i++ {
		(*players)[i%5].Draw(&deck)
	}
}
