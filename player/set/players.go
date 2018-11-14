package set

import (
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Players struct
type Players []*player.Player

// Add func
func (set *Players) Add(p player.Player) {
	*set = append(*set, &p)
}
