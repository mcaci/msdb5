package board

import (
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Players func
func (b *Board) Players() player.Players {
	return b.players
}
