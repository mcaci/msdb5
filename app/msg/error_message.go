package msg

import (
	"fmt"

	"github.com/mcaci/msdb5/app/phase"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Error func
func Error(text string, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	msg := printer.Sprintf(text)
	return fmt.Errorf(msg)
}

func unexpectedPlayerErr(name string, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	//g.CurrentPlayer().Name()
	msg := printer.Sprintf("Expecting player %s to play", name)
	return fmt.Errorf(msg)
}

func unexpectedPhaseErr(input, current phase.ID, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	msg := printer.Sprintf("Phase is not %d but %d", input, current)
	return fmt.Errorf(msg)
}
