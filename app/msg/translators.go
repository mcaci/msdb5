package msg

import (
	"strings"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"golang.org/x/text/message"
)

// TranslateCard func
func TranslateCard(c card.ID, printer *message.Printer) string {
	seeds := []string{printer.Sprintf("Coin"), printer.Sprintf("Cup"),
		printer.Sprintf("Sword"), printer.Sprintf("Cudgel")}
	return printer.Sprintf("(%d of %s)", c.Number(), seeds[c.Seed()])
}

// TranslateCards func
func TranslateCards(cards deck.Cards, printer *message.Printer) string {
	mappedCards := make([]string, 0, len(cards))
	for _, c := range cards {
		mappedCards = append(mappedCards, TranslateCard(c, printer))
	}
	return strings.Join(mappedCards, ",")
}
