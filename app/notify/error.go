package notify

import (
	"io"

	"golang.org/x/text/message"
)

// Err func
func Err(to io.Writer, err error, gameInfo senderInformer, rq requester) {
	printer := message.NewPrinter(gameInfo.Lang())
	ErrToConsole(to, err, gameInfo, rq)
	write(gameInfo.Sender(rq.From()), printer.Sprintf("Error: %+v\n", err))
}
