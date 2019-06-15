package gamelog

import (
	"fmt"
	"io"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// NotifyError func
func NotifyError(err error, gameInfo senderInformer, rq requester, notify func(*player.Player, string), to io.Writer) {
	sender := gameInfo.Sender(rq.From())
	ErrToConsole(sender.Name(), rq.Action(), err, to)
	errToSender(sender, err, notify)
}

func errToSender(sender *player.Player, err error, notify func(*player.Player, string)) {
	notify(sender, fmt.Sprintf("Error: %+v\n", err))
}

// ErrToConsole func
func ErrToConsole(senderName, request string, err error, to io.Writer) {
	msg := fmt.Sprintf("New Action by %s: %s\n"+
		"Error raised: %+v\n",
		senderName, request, err)
	write(to, msg)
}
