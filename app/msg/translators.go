package msg

import (
	"sort"
	"strings"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/briscola"
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
}

// TranslateCards func
func TranslateCards(cards set.Cards, printer *message.Printer) string {
	mCards := mappedCards(cards, printer)
	return strings.Join(mCards, ",")
}

// TranslateHand func
func TranslateHand(cards set.Cards, br *card.Seed, printer *message.Printer) string {
	mCards := mappedCards(cards, printer)
	sort.Sort(briscola.NewSorted(cards, br))
	return strings.Join(mCards, ",")
}

type statusProvider interface {
	AuctionScore() *auction.Score
	Briscola() card.Item
	CurrentPlayer() *player.Player
	Phase() phase.ID
	PlayedCard() *card.Item
	PlayedCards() *set.Cards
}

// TranslateGameStatus func
func TranslateGameStatus(g statusProvider, printer *message.Printer) string {
	return printer.Sprintf("Game: (Turn of: %s, Companion is: %s, Played cards: %s, Last card: %s, Auction score: %d, Phase: %s)",
		g.CurrentPlayer().Name(), TranslateCard(g.Briscola(), printer),
		TranslateCards(*g.PlayedCards(), printer), TranslateCard(*g.PlayedCard(), printer),
		*g.AuctionScore(), TranslatePhase(g.Phase(), printer))
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
	Briscola() *card.Seed
	Phase() phase.ID
	SideDeck() *set.Cards
}

// TranslatePlayer func
func TranslatePlayer(gameInfo selfInformer, pl *player.Player, printer *message.Printer) string {
	return printer.Sprintf("Player: (Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)",
		pl.Name(), TranslateHand(*pl.Hand(), &gameInfo.Briscola(), printer), TranslateCards(*pl.Pile(), printer), player.Folded(pl))
}

// TranslateSideDeck func
func TranslateSideDeck(gameInfo selfInformer, pl *player.Player, printer *message.Printer) string {
	return printer.Sprintf("Side deck: %s\n", TranslateCards(*gameInfo.SideDeck(), printer))
}
