package game

import (
	"fmt"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type playerPredicate func(p *player.Player) bool

func findCriteria(g *Game, request, origin string) playerPredicate {
	var expectedPlayerFinder playerPredicate
	action := strings.Split(request, "#")[0]
	switch action {
	case "Join":
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsNameEmpty() }
	case "Origin":
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsSameHost(origin) }
	default:
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsExpectedPlayer(g.CurrentPlayer(), origin) }
	}
	return expectedPlayerFinder
}

func verifyPlayer(g *Game, request, origin string, notify func(*player.Player, string)) error {
	criteria := findCriteria(g, request, origin)
	_, actingPlayer, err := g.players.Find(criteria)
	if err != nil {
		err = fmt.Errorf("%v. Expecting player %s to play", err, g.CurrentPlayer().Name())
		gamelog.ToConsole(g, g.sender(origin), request, err)
		notify(g.sender(origin), err.Error())
		return err
	}
	if g.CurrentPlayer() == actingPlayer {
		return nil
	}
	trackActing(&g.lastPlaying, actingPlayer)
	return nil
}

func verifyPhase(g *Game, request, origin string, notify func(*player.Player, string)) error {
	currentPhase := g.phase
	inputPhase, err := phase.ToID(request)
	if err != nil {
		gamelog.ToConsole(g, g.sender(origin), request, err)
		notify(g.sender(origin), err.Error())
		return err
	}
	if currentPhase != inputPhase {
		err = fmt.Errorf("Phase is not %d but %d", inputPhase, currentPhase)
		gamelog.ToConsole(g, g.sender(origin), request, err)
		notify(g.sender(origin), err.Error())
		return err
	}
	return nil
}
