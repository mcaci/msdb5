package briscola

import "github.com/mcaci/ita-cards/set"

type Deck struct{ set.Cards }

func NewDeck() *Deck {
	return &Deck{Cards: set.Deck()}
}

type PlayedCards struct{ *set.Cards }

func (c PlayedCards) Pile() *set.Cards {
	if len(*c.Cards) == nPlayers {
		return (*set.Cards)(c.Cards)
	}
	return &set.Cards{}
}
