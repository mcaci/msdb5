package msg

import (
	"fmt"

	"github.com/mcaci/msdb5/app/phase"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// UnexpectedPlayerErr func
func UnexpectedPlayerErr(name string, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	msg := printer.Sprintf("Expecting player %s to play", name)
	return fmt.Errorf(msg)
}

// UnexpectedPhaseErr func
func UnexpectedPhaseErr(input, current phase.ID, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	msg := printer.Sprintf("Phase is not %d but %d", input, current)
	return fmt.Errorf(msg)
}
