package gamelog

import (
	"io"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type senderInformer interface {
	Sender(string) *player.Player
}

type requester interface {
	From() string
	Action() string
}

// ToConsole func
func ToConsole(to io.Writer, gameInfo senderInformer, rq requester) {
	sender := gameInfo.Sender(rq.From())
	printer := message.NewPrinter(language.English)
	msg := printer.Sprintf("New Action by %s: %s\n"+
		"Sender info: %+v\n"+
		"Game info: %+v\n",
		sender.Name(), rq.Action(), sender, gameInfo)
	write(to, msg)
}
