package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Players func
func (b *Board) Players() player.Players {
	return b.players
}

// PChans func
func (b *Board) PChans() []chan card.ID {
	return b.pChans
}
