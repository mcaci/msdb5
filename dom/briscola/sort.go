package briscola

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

type sortedCard struct {
	cards    set.Cards
	briscola *card.Seed
}

func (ids sortedCard) Len() int { return len(ids.cards) }

func (ids sortedCard) Less(i, j int) bool {
	if ids.briscola == nil {
		return isOtherHigher(ids.cards[i], ids.cards[j])
	}
	return doesOtherCardWin(ids.cards[i], ids.cards[j], *ids.briscola)
}

func (ids sortedCard) Swap(i, j int) { ids.cards[i], ids.cards[j] = ids.cards[j], ids.cards[i] }
