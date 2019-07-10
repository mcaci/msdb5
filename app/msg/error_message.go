package msg

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// ErrPlayerNotFound func
func ErrPlayerNotFound(name string, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	msg := printer.Sprintf("Expecting player %s to play", name)
	return fmt.Errorf(msg)
}

// ErrPhaseNotExpected func
func ErrPhaseNotExpected(inputPhase, currentPhase uint8, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	msg := printer.Sprintf("Phase is not %d but %d", inputPhase, currentPhase)
	return fmt.Errorf(msg)
}

// ErrInvalidAction func
func ErrInvalidAction(action string, lang language.Tag) error {
	printer := message.NewPrinter(lang)
	msg := printer.Sprintf("Action %s not valid", action)
	return fmt.Errorf(msg)
}
