package msg

import (
	"testing"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/language"
)

type selfStruct struct{ ph phase.ID }

func (selfStruct) Lang() language.Tag    { return language.English }
func (s selfStruct) Phase() phase.ID     { return s.ph }
func (selfStruct) SideDeck() *deck.Cards { return &deck.Cards{1} }

func TestValidMessage(t *testing.T) {
	str := CreateInGameMsg(selfStruct{0}, player.New())
	if str == "" {
		t.Fatal("Expecting a message")
	}
}

func TestValidMessagePlusExchangeInfo(t *testing.T) {
	str := CreateInGameMsg(selfStruct{phase.ExchangingCards}, player.New())
	if str == "" {
		t.Fatal("Expecting a message")
	}
}
