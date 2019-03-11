package orchestrator

import "github.com/nikiforosFreespirit/msdb5/player"

var isNameEmpty = func(p *player.Player) bool { return p.IsName("") }
var isHandEmpty = func(p *player.Player) bool { return len(*p.Hand()) == 0 }
var folded = func(p *player.Player) bool { return p.Folded() }
var notFolded = func(p *player.Player) bool { return !p.Folded() }

var isRemoteHost = func(p *player.Player, origin string) bool { return p.IsRemoteHost(origin) }
var isInTurn = func(p *player.Player, g *Game) bool { return p == g.players[g.playerInTurn] }
var isExpectedPlayer = func(p *player.Player, g *Game, origin string) bool {
	return isInTurn(p, g) && isRemoteHost(p, origin)
}
