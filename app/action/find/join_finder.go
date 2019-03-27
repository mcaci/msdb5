package find

import (
	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/game"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type JoinFinder game.Phase

func NewJoinFinder() action.Finder {
	return JoinFinder(game.Joining)
}

func (jf JoinFinder) Find(p *player.Player) bool { return p.IsNameEmpty() }
