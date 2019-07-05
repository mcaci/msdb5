package notify

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
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
	mappedCards := make([]string, 0)
	for _, c := range cards {
		mappedCards = append(mappedCards, TranslateCard(c, printer))
	}
	return strings.Join(mappedCards, ",")
}

// TranslatePlayer func
func TranslatePlayer(pl player.Player, printer *message.Printer) string {
	return printer.Sprintf("(Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)",
		pl.Name(), TranslateCards(*pl.Hand(), printer), TranslateCards(*pl.Pile(), printer), pl.Folded())
}
