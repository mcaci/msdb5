package gamelog

import (
	"fmt"
	"io"
)

// NotifyError func
func NotifyError(dest io.Writer, err error, gameInfo senderInformer, rq requester) {
	ErrToConsole(dest, err, gameInfo, rq)
	write(gameInfo.Sender(rq.From()), fmt.Sprintf("Error: %+v\n", err))
}

// ErrToConsole func
func ErrToConsole(dest io.Writer, err error, gameInfo senderInformer, rq requester) {
	sender := gameInfo.Sender(rq.From())
	write(dest, fmt.Sprintf("New Action by %s: %s\nError raised: %+v\n", sender.Name(), rq.Action(), err))
}
