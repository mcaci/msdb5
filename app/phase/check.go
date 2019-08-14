package phase

import (
	"errors"
)

// ErrUnexpectedPhase error
var ErrUnexpectedPhase = errors.New("Unexpected phase")

func Check(g interface{ Phase() ID }, rq interface{ Action() string }) error {
	inputPhase, err := ToID(rq)
	if err == nil && g.Phase() != inputPhase {
		err = ErrUnexpectedPhase
	}
	return err
}
