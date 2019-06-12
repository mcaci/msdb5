package game

import (
	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func notifyPlayer(g *Game, rq *req, notify func(*player.Player, string)) error {
	notify(g.LastPlayer(), gamelog.ToMe(g))
	return nil
}

func notifyAll(g *Game, rq *req, notify func(*player.Player, string)) error {
	for _, pl := range g.players {
		notify(pl, gamelog.ToAll(g))
	}
	return nil
}
