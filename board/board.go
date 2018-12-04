package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	cset "github.com/nikiforosFreespirit/msdb5/card/set"
	"github.com/nikiforosFreespirit/msdb5/player"
	pset "github.com/nikiforosFreespirit/msdb5/player/set"
)

// Board struct
type Board struct {
	players      pset.Players
	pChans       []chan card.ID
	playedCards  cset.Cards
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

func makePlayers(b *Board) {
	for i := 0; i < 5; i++ {
		b.players = append(b.players, player.New())
	}
}

func playersDrawAllCards(players *pset.Players) {
	deck := cset.Deck()
	for i := 0; i < cset.DeckSize; i++ {
		(*players)[i%5].Draw(&deck)
	}
}
