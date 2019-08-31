package msg

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/language"
)

type selfStruct struct{ ph phase.ID }

func (s selfStruct) Phase() phase.ID    { return s.ph }
func (selfStruct) SideDeck() *set.Cards { return set.NewMust(1) }

func TestValidMessage(t *testing.T) {
	str := CreateInGameMsg(selfStruct{0}, player.New(), language.English)
	if str == "" {
		t.Fatal("Expecting a message")
	}
}

func TestValidMessagePlusExchangeInfo(t *testing.T) {
	str := CreateInGameMsg(selfStruct{phase.ExchangingCards}, player.New(), language.English)
	if str == "" {
		t.Fatal("Expecting a message")
	}
}
