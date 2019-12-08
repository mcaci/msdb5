package cardsort

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

type SortedCard struct {
	cards    set.Cards
	briscola *card.Seed
}

func NewSorted(cards set.Cards, b *card.Seed) *SortedCard {
	return &SortedCard{cards, b}
}

func (ids SortedCard) Len() int { return len(ids.cards) }

func (ids SortedCard) Less(i, j int) bool {
	all := []uint8{1, 3, 10, 9, 8, 7, 6, 5, 4, 2,
		11, 13, 20, 19, 18, 17, 16, 15, 14, 12,
		21, 23, 30, 29, 28, 27, 26, 25, 24, 22,
		31, 33, 40, 39, 38, 37, 36, 35, 34, 32}
	iID, jID := ids.cards[i].ToID(), ids.cards[j].ToID()
	var iIdx, jIdx int
	for idx, cID := range all {
		if iID == cID {
			iIdx = idx
		}
		if jID == cID {
			jIdx = idx
		}
	}
	return iIdx <= jIdx
}

func (ids SortedCard) Swap(i, j int) { ids.cards[i], ids.cards[j] = ids.cards[j], ids.cards[i] }
