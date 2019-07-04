package notify

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"golang.org/x/text/message"
)

// TranslateCards func
func TranslateCards(cards deck.Cards, printer *message.Printer) string {
	seeds := []string{printer.Sprint("Coin"), printer.Sprint("Cup"),
		printer.Sprint("Sword"), printer.Sprint("Cudgel")}
	mappedCards := make([]string, 0)
	for _, c := range cards {
		card := printer.Sprintf("(%d of %s)", c.Number(), seeds[c.Seed()])
		mappedCards = append(mappedCards, card)
	}
	return strings.Join(mappedCards, ",")
}

// TranslatePlayer func
func TranslatePlayer(pl player.Player, printer *message.Printer) string {
	return printer.Sprintf("(Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)",
		pl.Name(), TranslateCards(*pl.Hand(), printer), TranslateCards(*pl.Pile(), printer), pl.Folded())
}
