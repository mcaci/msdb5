package orchestrator

import "github.com/nikiforosFreespirit/msdb5/player"

var isNameEmpty = func(p *player.Player) bool { return p.IsName("") }
var isHandEmpty = func(p *player.Player) bool { return len(*p.Hand()) == 0 }
var folded = func(p *player.Player) bool { return p.Folded() }
var notFolded = func(p *player.Player) bool { return !p.Folded() }

var isSameHost = func(p *player.Player, origin string) bool { return p.IsSameHost(origin) }
var isSame = func(p, other *player.Player) bool { return p == other }
var isExpectedPlayer = func(p *player.Player, g *Game, origin string) bool {
	return isSame(p, g.players[g.playerInTurn]) && isSameHost(p, origin)
}
