package msg

import (
	"sort"
	"strings"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/msg/cardsort"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/message"
)

// TranslateCard func
func TranslateCard(c card.Item, printer *message.Printer) string {
	if c.Number() == 0 {
		return printer.Sprintf("(Undefined card)")
	}
	seeds := []string{printer.Sprintf("Coin"), printer.Sprintf("Cup"),
		printer.Sprintf("Sword"), printer.Sprintf("Cudgel")}
	return printer.Sprintf("(%d of %s)", c.Number(), seeds[c.Seed()])
}

func mappedCards(cards set.Cards, printer *message.Printer) []string {
	mCards := make([]string, 0, len(cards))
	for _, c := range cards {
		mCards = append(mCards, TranslateCard(c, printer))
	}
	return mCards
}

// TranslateCards func
func TranslateCards(cards set.Cards, printer *message.Printer) string {
	mCards := mappedCards(cards, printer)
	return strings.Join(mCards, ",")
}

// TranslateHand func
func TranslateHand(cards set.Cards, br *card.Seed, printer *message.Printer) string {
	sort.Sort(cardsort.NewSorted(cards, br))
	return TranslateCards(cards, printer)
}

// TranslateSideDeck func
func TranslateSideDeck(gameInfo interface{ SideDeck() *set.Cards }, pl *player.Player, printer *message.Printer) string {
	return printer.Sprintf("Side deck: %s\n", TranslateCards(*gameInfo.SideDeck(), printer))
}
