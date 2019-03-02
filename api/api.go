package api

import "github.com/nikiforosFreespirit/msdb5/display"

// Action interface
type Action interface {
	Action(request, origin string) (display.Info, display.Info, error)
}
