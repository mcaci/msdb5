package orchestrator

import (
	"log"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Action func
func (g *Game) Action(request, origin string) (all []display.Info, me []display.Info, err error) {
	action := strings.Split(string(request), "#")[0]
	all, me, err = g.actionMap[action](request, origin)
	logEndRound(g, request, origin, err)
	return
}

func logEndRound(g *Game, request, origin string, err error) {
	playerLogged, _ := g.players.Find(func(p *player.Player) bool { return isRemoteHost(p, origin) })
	log.Printf("New Action by %s\n", playerLogged.Name().Display())
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %v\n", err)
	log.Printf("Game info after action: %s\n", g.String())
}
