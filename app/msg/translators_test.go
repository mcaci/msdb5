package msg

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
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
