package msg

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/message"
)

type statusProvider interface {
	AuctionScore() *auction.Score
	Briscola() card.Item
	CurrentPlayer() *player.Player
	Phase() phase.ID
	PlayedCard() card.Item
	PlayedCards() *set.Cards
}

// TranslateGameStatus func
func TranslateGameStatus(g statusProvider, printer *message.Printer) string {
	var c card.Item
	if g.Phase() == phase.PlayingCards {
		c = g.PlayedCard()
	}
	return printer.Sprintf("Game: (Turn of: %s, Companion is: %s, Played cards: %s, Last card: %s, Auction score: %d, Phase: %s)\n",
		g.CurrentPlayer().Name(), TranslateCard(g.Briscola(), printer),
		TranslateCards(*g.PlayedCards(), printer), TranslateCard(c, printer),
		*g.AuctionScore(), TranslatePhase(g.Phase(), printer))
}

// TranslatePhase func
func TranslatePhase(p phase.ID, printer *message.Printer) string {
	phases := []string{printer.Sprintf("Join"), printer.Sprintf("Auction"),
		printer.Sprintf("Exchange"), printer.Sprintf("Companion"),
		printer.Sprintf("Card"), printer.Sprintf("End")}
	return phases[p]
}
