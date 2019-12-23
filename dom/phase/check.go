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

type Info struct {
	phase  ID
	action string
}

func NewInfo(phase ID, action string) Info {
	return Info{phase, action}
}

func (s Info) Action() string { return s.action }
func (s Info) Phase() ID      { return s.phase }

func Check(g phaseChecker) error {
	inputPhase, err := ToID(g.Action())
	if err == nil && g.Phase() != inputPhase {
		err = ErrUnexpectedPhase
	}
	return err
}
