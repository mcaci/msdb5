package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
)

// Action func
func (g *Game) Action(request, origin string) (all []display.Info, me []display.Info, err error) {
	// logBeforeRound(g, request, origin)
	action := strings.Split(string(request), "#")[0]
	all, me, err = g.actionMap[action](request, origin)
	// logEndRound(g, request, origin, err)
	if g.phase == end {
		all, me, err = endGame(g)
	}
	return
}
