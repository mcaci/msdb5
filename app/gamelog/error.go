package gamelog

import (
	"io"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// NotifyError func
func NotifyError(dest io.Writer, err error, gameInfo senderInformer, rq requester) {
	printer := message.NewPrinter(language.English)
	ErrToConsole(dest, err, gameInfo, rq)
	write(gameInfo.Sender(rq.From()), printer.Sprintf("Error: %+v\n", err))
}

// ErrToConsole func
func ErrToConsole(dest io.Writer, err error, gameInfo senderInformer, rq requester) {
	printer := message.NewPrinter(language.English)
	sender := gameInfo.Sender(rq.From())
	write(dest, printer.Sprintf("New Action by %s: %s\nError raised: %+v\n", sender.Name(), rq.Action(), err))
}
