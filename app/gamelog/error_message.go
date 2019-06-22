package gamelog

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// ErrPlayerNotFound func
func ErrPlayerNotFound(err error, name string) error {
	printer := message.NewPrinter(language.English)
	msg := printer.Sprintf("%v. Expecting player %s to play", err, name)
	return fmt.Errorf(msg)
}

// ErrPhaseNotExpected func
func ErrPhaseNotExpected(inputPhase, currentPhase uint8) error {
	printer := message.NewPrinter(language.English)
	msg := printer.Sprintf("Phase is not %d but %d", inputPhase, currentPhase)
	return fmt.Errorf(msg)
}

// ErrInvalidAction func
func ErrInvalidAction(action string) error {
	printer := message.NewPrinter(language.English)
	msg := printer.Sprintf("Action %s not valid", action)
	return fmt.Errorf(msg)
}
