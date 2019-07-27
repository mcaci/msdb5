package msg

import (
	"strings"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/message"
)

// TranslateCard func
func TranslateCard(c card.ID, printer *message.Printer) string {
	if c == 0 {
		return printer.Sprintf("(Undefined card)")
	}
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

type statusProvider interface {
	AuctionScore() *auction.Score
	Briscola() card.ID
	CurrentPlayer() *player.Player
	Phase() phase.ID
	PlayedCards() *deck.Cards
}

// TranslateGameStatus func
func TranslateGameStatus(g statusProvider, printer *message.Printer) string {
	return printer.Sprintf("(Turn of: %s, Companion is: %s, Played cards: %s, Auction score: %d, Phase: %s)",
		g.CurrentPlayer().Name(), TranslateCard(g.Briscola(), printer), TranslateCards(*g.PlayedCards(), printer), g.AuctionScore(), g.Phase())
}
