package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Cards are slices of card.Cards
type Cards []card.Card

// Board struct
type Board struct {
	deck        deck.Deck
	players     []player.Player
	playedCards *Cards
}

// New func
func New() *Board {
	var b Board

	b.deck = deck.New()

	b.players = make([]player.Player, 5)
	for i := range b.players {
		b.players[i] = player.New()
	}
	for i := 0; !b.deck.IsEmpty(); i++ {
		b.players[i%5].Draw(b.deck)
	}

	b.playedCards = new(Cards)

	return &b
}

// Deck func
func (b *Board) Deck() deck.Deck {
	return b.deck
}

// Players func
func (b *Board) Players() []player.Player {
	return b.players
}

// PlayedCards func
func (b *Board) PlayedCards() *Cards {
	return b.playedCards
}

// Add func
func (cards *Cards) Add(c card.Card) {
	*cards = append(*cards, c)
}

// Has func
func (cards Cards) Has(c card.Card) bool {
	var cardFound bool
	for _, card := range cards {
		cardFound = (c == card)
		if cardFound {
			break
		}
	}
	return cardFound
}
