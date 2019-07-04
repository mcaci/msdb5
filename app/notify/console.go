package notify

import (
	"fmt"
	"io"

	"golang.org/x/text/language"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type senderInformer interface {
	Sender(string) *player.Player
	Lang() language.Tag
}

type requester interface {
	From() string
	Action() string
}

func sendToConsole(to io.Writer, senderName, action, addInfo string) {
	write(to, fmt.Sprintf("New Action by %s: %s\n%s\n", senderName, action, addInfo))
}

// ToConsole func
func ToConsole(to io.Writer, gameInfo senderInformer, rq requester) {
	sender := gameInfo.Sender(rq.From())
	addInfo := fmt.Sprintf("Sender info: %+v\nGame info: %+v", sender, gameInfo)
	sendToConsole(to, sender.Name(), rq.Action(), addInfo)
}

// ErrToConsole func
func ErrToConsole(to io.Writer, err error, gameInfo senderInformer, rq requester) {
	sender := gameInfo.Sender(rq.From())
	addInfo := fmt.Sprintf("Error raised: %+v", err)
	sendToConsole(to, sender.Name(), rq.Action(), addInfo)
}
