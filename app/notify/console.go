package notify

import (
	"io"

	"golang.org/x/text/language"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"golang.org/x/text/message"
)

type senderInformer interface {
	Sender(string) *player.Player
	Lang() language.Tag
}

type requester interface {
	From() string
	Action() string
}

// ToConsole func
func ToConsole(to io.Writer, gameInfo senderInformer, rq requester) {
	sender := gameInfo.Sender(rq.From())
	printer := message.NewPrinter(gameInfo.Lang())
	msg := printer.Sprintf("New Action by %s: %s\n"+
		"Sender info: %+v\n"+
		"Game info: %+v\n",
		sender.Name(), rq.Action(), sender, gameInfo)
	write(to, msg)
}
