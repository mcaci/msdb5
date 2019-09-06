package msg

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type msgToPlTest struct{}

func (msgToPlTest) Briscola() card.Item { return *card.MustID(1) }

func TestValidMessage(t *testing.T) {
	printer := message.NewPrinter(language.English)
	str := TranslatePlayer(player.New(), msgToPlTest{}, printer)
	if str == "" {
		t.Fatal("Expecting a message")
	}
}
