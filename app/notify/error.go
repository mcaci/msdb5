package notify

import (
	"fmt"
	"io"

	"golang.org/x/text/message"
)

// Err func
func Err(dest io.Writer, err error, gameInfo senderInformer, rq requester) {
	printer := message.NewPrinter(gameInfo.Lang())
	ErrToConsole(dest, err, gameInfo, rq)
	write(gameInfo.Sender(rq.From()), printer.Sprintf("Error: %+v\n", err))
}

// ErrToConsole func
func ErrToConsole(dest io.Writer, err error, gameInfo senderInformer, rq requester) {
	sender := gameInfo.Sender(rq.From())
	write(dest, fmt.Sprintf("New Action by %s: %s\nError raised: %+v\n", sender.Name(), rq.Action(), err))
}
