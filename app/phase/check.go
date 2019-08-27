package phase

import (
	"errors"
)

// ErrUnexpectedPhase error
var ErrUnexpectedPhase = errors.New("Unexpected phase")

type phaseChecker interface {
	Phase() ID
	Action() string
}

func Check(g phaseChecker) error {
	inputPhase, err := ToID(g.Action())
	if err == nil && g.Phase() != inputPhase {
		err = ErrUnexpectedPhase
	}
	return err
}
