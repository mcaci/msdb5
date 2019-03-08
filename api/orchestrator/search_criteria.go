package orchestrator

import "github.com/nikiforosFreespirit/msdb5/player"

var isNameEmpty = func(p *player.Player) bool { return p.IsName("") }

var isRemoteHost = func(p *player.Player, origin string) bool { return p.IsRemoteHost(origin) }

var notFolded = func(p *player.Player) bool { return !p.Folded() }

var folded = func(p *player.Player) bool { return p.Folded() }

var isActive = func(g *Game, p *player.Player, origin string) bool {
	return isRemoteHost(p, origin) && p == g.players[g.playerInTurn]
}
