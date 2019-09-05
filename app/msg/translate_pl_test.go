package msg

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type selfStruct struct{ ph phase.ID }

func (s selfStruct) Phase() phase.ID    { return s.ph }
func (selfStruct) SideDeck() *set.Cards { return set.NewMust(1) }

func TestValidMessage(t *testing.T) {
	printer := message.NewPrinter(language.English)
	str := TranslatePlayer(selfStruct{0}, player.New(), printer)
	if str == "" {
		t.Fatal("Expecting a message")
	}
}

func TestValidMessagePlusExchangeInfo(t *testing.T) {
	printer := message.NewPrinter(language.English)
	str := TranslatePlayer(selfStruct{phase.ExchangingCards}, player.New(), printer)
	if str == "" {
		t.Fatal("Expecting a message")
	}
}
