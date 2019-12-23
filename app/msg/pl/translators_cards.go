package pl

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/msg/cardsort"
	"golang.org/x/text/message"
)

func translateCard(c card.Item, printer *message.Printer) string {
	if c.Number() == 0 {
		return undefinedCard(printer)
	}
	return fmt.Sprintf("(%d of %s)", c.Number(), seeds(printer, uint8(c.Seed())))
}

func translateCards(cards set.Cards, printer *message.Printer) string {
	mCards := make([]string, 0, len(cards))
	for _, c := range cards {
		mCards = append(mCards, translateCard(c, printer))
	}
	return strings.Join(mCards, ",")
}

func translateHand(cards set.Cards, br *card.Seed, printer *message.Printer) string {
	sort.Sort(cardsort.NewSorted(cards, br))
	return translateCards(cards, printer)
}
