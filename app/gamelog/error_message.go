package gamelog

import (
	"fmt"
)

// ErrPlayerNotFound func
func ErrPlayerNotFound(err error, name string) error {
	return fmt.Errorf("%v. Expecting player %s to play", err, name)
}

// ErrPhaseNotExpected func
func ErrPhaseNotExpected(inputPhase, currentPhase uint8) error {
	return fmt.Errorf("Phase is not %d but %d", inputPhase, currentPhase)
}

// ErrInvalidAction func
func ErrInvalidAction(action string) error {
	return fmt.Errorf("Action %s not valid", action)
}
