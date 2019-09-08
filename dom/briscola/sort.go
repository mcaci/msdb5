package briscola

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

type SortedCard struct {
	cards    set.Cards
	briscola *card.Seed
}

func NewSorted(cards set.Cards, briscola *card.Seed) *SortedCard {
	return &SortedCard{cards, briscola}
}

func (ids SortedCard) Len() int { return len(ids.cards) }

func (ids SortedCard) Less(i, j int) bool {
	if ids.briscola == nil {
		return !isOtherHigher(ids.cards[i], ids.cards[j])
	}
	return !doesOtherCardWin(ids.cards[i], ids.cards[j], *ids.briscola)
}

func (ids SortedCard) Swap(i, j int) { ids.cards[i], ids.cards[j] = ids.cards[j], ids.cards[i] }
