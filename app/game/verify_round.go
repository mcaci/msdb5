package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type playerPredicate func(p *player.Player) bool

func findCriteria(g *Game, rq *req) playerPredicate {
	var expectedPlayerFinder playerPredicate
	switch rq.Action() {
	case "Join":
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsNameEmpty() }
	case "Origin":
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsSameHost(rq.From()) }
	default:
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsExpectedPlayer(g.CurrentPlayer(), rq.From()) }
	}
	return expectedPlayerFinder
}

func verifyPlayer(g *Game, rq *req, notify func(*player.Player, string)) error {
	criteria := findCriteria(g, rq)
	_, actingPlayer, err := g.players.Find(criteria)
	if err != nil {
		err = fmt.Errorf("%v. Expecting player %s to play", err, g.CurrentPlayer().Name())
		sender := g.sender(rq.From())
		gamelog.ToConsole(g, sender, rq.Action(), err)
		notify(sender, err.Error())
		return err
	}
	if g.CurrentPlayer() == actingPlayer {
		return nil
	}
	trackActing(&g.lastPlaying, actingPlayer)
	return nil
}

func verifyPhase(g *Game, rq *req, notify func(*player.Player, string)) error {
	currentPhase := g.phase
	inputPhase, err := phase.ToID(rq.Action())
	if err == nil && currentPhase == inputPhase {
		return nil
	}
	sender := g.sender(rq.From())
	if err == nil && currentPhase != inputPhase {
		err = fmt.Errorf("Phase is not %d but %d", inputPhase, currentPhase)
	}
	gamelog.ToConsole(g, sender, rq.Action(), err)
	notify(sender, err.Error())
	return err
}
