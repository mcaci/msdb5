package msg

import (
	"strings"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
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

// TranslateCards func
func TranslateCards(cards set.Cards, printer *message.Printer) string {
	mappedCards := make([]string, 0, len(cards))
	for _, c := range cards {
		mappedCards = append(mappedCards, TranslateCard(c, printer))
	}
	return strings.Join(mappedCards, ",")
}

type statusProvider interface {
	AuctionScore() *auction.Score
	Briscola() card.Item
	CurrentPlayer() *player.Player
	Phase() phase.ID
	PlayedCards() *set.Cards
}

// TranslateGameStatus func
func TranslateGameStatus(g statusProvider, printer *message.Printer) string {
	return printer.Sprintf("Game: (Turn of: %s, Companion is: %s, Played cards: %s, Auction score: %d, Phase: %s)",
		g.CurrentPlayer().Name(), TranslateCard(g.Briscola(), printer),
		TranslateCards(*g.PlayedCards(), printer), *g.AuctionScore(), TranslatePhase(g.Phase(), printer))
}

// TranslatePhase func
func TranslatePhase(p phase.ID, printer *message.Printer) string {
	phases := []string{printer.Sprintf("Join"), printer.Sprintf("Auction"),
		printer.Sprintf("Exchange"), printer.Sprintf("Companion"),
		printer.Sprintf("Card"), printer.Sprintf("End")}
	return phases[p]
}

type callersProvider interface {
	Caller() *player.Player
	Companion() *player.Player
}

// TranslateTeam func
func TranslateTeam(p *player.Player, g callersProvider, printer *message.Printer) string {
	team := printer.Sprintf("Callers")
	if p != g.Caller() && p != g.Companion() {
		team = printer.Sprintf("Others")
	}
	return printer.Sprintf("The end - %s team has all briscola cards", team)
}

type selfInformer interface {
	Phase() phase.ID
	SideDeck() *set.Cards
}

// TranslatePlayer func
func TranslatePlayer(gameInfo selfInformer, pl *player.Player, printer *message.Printer) string {
	me := printer.Sprintf("Player: (Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)",
		pl.Name(), TranslateCards(*pl.Hand(), printer), TranslateCards(*pl.Pile(), printer), player.Folded(pl))
	if gameInfo.Phase() == phase.ExchangingCards {
		me += " " + printer.Sprintf("Side deck: %s\n", TranslateCards(*gameInfo.SideDeck(), printer))
	}
	return me
}
