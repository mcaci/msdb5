package msg

import (
	"testing"

	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestValidMessage(t *testing.T) {
	printer := message.NewPrinter(language.English)
	str := TranslatePlayer(player.New(), printer)
	if str == "" {
		t.Fatal("Expecting a message")
	}
}
