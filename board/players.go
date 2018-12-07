package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player/set"
)

// Players func
func (b *Board) Players() set.Players {
	return b.players
}

// PChans func
func (b *Board) PChans() []chan card.ID {
	return b.pChans
}
