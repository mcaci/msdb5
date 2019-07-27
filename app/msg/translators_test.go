package msg

import (
	"testing"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestValidCardTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	c := card.ID(1)
	str := TranslateCard(c, printer)
	if str == "(Undefined card)" {
		t.Fatal("Expecting the translation of the card")
	}
}

func TestInvalidCardTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	var c card.ID
	str := TranslateCard(c, printer)
	if str != "(Undefined card)" {
		t.Fatal("Expecting the translation of the card")
	}
}

func TestValidCardsTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	c := deck.Cards{1, 2, 3, 4, 7}
	str := TranslateCards(c, printer)
	if str == "" {
		t.Fatal("Expecting the translation of the set of card")
	}
}

type fakeGameStatus struct{}

func (g fakeGameStatus) AuctionScore() *auction.Score  { v := auction.Score(1); return &v }
func (g fakeGameStatus) Briscola() card.ID             { return 1 }
func (g fakeGameStatus) CurrentPlayer() *player.Player { return player.New() }
func (g fakeGameStatus) Phase() phase.ID               { return 1 }
func (g fakeGameStatus) PlayedCards() *deck.Cards      { return &deck.Cards{1} }

func TestValidGameStatusTranslation(t *testing.T) {
	printer := message.NewPrinter(language.English)
	g := fakeGameStatus{}
	str := TranslateGameStatus(g, printer)
	if str == "" {
		t.Fatal("Expecting the translation of the set of card")
	}
}
