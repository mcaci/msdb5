package gamelog

import (
	"fmt"
	"io"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
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
	msg := fmt.Sprintf("New Action by %s: %s\n"+
		"Sender info: %+v\n"+
		"Game info: %+v\n",
		sender.Name(), rq.Action(), sender, gameInfo)
	write(to, msg)
}
