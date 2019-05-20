package find

import (
	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type JoinFinder phase.ID

func NewJoinFinder() action.Finder {
	return JoinFinder(phase.Joining)
}

func (jf JoinFinder) Find(p *player.Player) bool { return p.IsNameEmpty() }
