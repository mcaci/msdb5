package find

import (
	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
)

type JoinFinder game.Phase

func NewJoinFinder() action.Finder {
	return JoinFinder(game.Joining)
}

func (jf JoinFinder) Find(p *player.Player) bool { return p.IsNameEmpty() }
