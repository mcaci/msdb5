package board

import (
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Players func
func (b *Board) Players() []*player.Player {
	return b.players
}
