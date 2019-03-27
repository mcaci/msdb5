package find

import (
	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/player"
)

type PlayerFinder struct {
	origin       string
	playerInTurn *player.Player
}

func NewPlayerFinder(origin string, playerInTurn *player.Player) action.Finder {
	return &PlayerFinder{origin, playerInTurn}
}

func (pf PlayerFinder) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(pf.playerInTurn, pf.origin)
}
