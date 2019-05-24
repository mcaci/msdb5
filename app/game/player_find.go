package game

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type playerPredicate func(p *player.Player) bool

func findCriteria(g *Game, request, origin string) playerPredicate {
	var expectedPlayerFinder playerPredicate
	action := strings.Split(request, "#")[0]
	switch action {
	case "Join":
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsNameEmpty() }
	default:
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsExpectedPlayer(g.CurrentPlayer(), origin) }
	}
	return expectedPlayerFinder
}
