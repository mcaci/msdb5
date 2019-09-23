package msg

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestValidCardTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	c := *card.MustID(1)
	str := TranslateCard(c, printer)
	if str == "(Undefined card)" {
		t.Fatal("Expecting the translation of the card")
	}
}

func TestInvalidCardTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	var c card.Item
	str := TranslateCard(c, printer)
	if str != "(Undefined card)" {
		t.Fatal("Expecting the translation of the card")
	}
}

func TestValidCardsTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	c := set.NewMust(1, 2, 3, 4, 7)
	str := TranslateCards(*c, printer)
	if str == "" {
		t.Fatal("Expecting the translation of the set of card")
	}
}

type fakeGameStatus struct{}

func (g fakeGameStatus) AuctionScore() *auction.Score  { v := auction.Score(1); return &v }
func (g fakeGameStatus) Briscola() card.Item           { return *card.MustID(1) }
func (g fakeGameStatus) CurrentPlayer() *player.Player { return player.New() }
func (g fakeGameStatus) Phase() phase.ID               { return 1 }
func (g fakeGameStatus) PlayedCard() card.Item         { return *card.MustID(2) }
func (g fakeGameStatus) PlayedCards() *set.Cards       { return set.NewMust(1) }

func TestValidGameStatusTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	g := fakeGameStatus{}
	str := TranslateGameStatus(g, printer)
	if str == "" {
		t.Fatal("Expecting the translation of the set of card")
	}
}

func TestValidPhaseTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	p := phase.Joining
	str := TranslatePhase(p, printer)
	if str == "" {
		t.Fatal("Expecting the translation of the phase")
	}
}
