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

func sendToConsole(to io.Writer, gameInfo senderInformer, rq requester, addInfo string) {
	sender := gameInfo.Sender(rq.From())
	write(to, fmt.Sprintf("New Action by %s: %s\n", sender.Name(), rq.Action())+addInfo)
}

// ToConsole func
func ToConsole(to io.Writer, gameInfo senderInformer, rq requester) {
	sender := gameInfo.Sender(rq.From())
	sendToConsole(to, gameInfo, rq, fmt.Sprintf("Sender info: %+v\nGame info: %+v\n", sender, gameInfo))
}

// ErrToConsole func
func ErrToConsole(to io.Writer, err error, gameInfo senderInformer, rq requester) {
	sendToConsole(to, gameInfo, rq, fmt.Sprintf("Error raised: %+v\n", err))
}
