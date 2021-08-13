package end

import (
	"github.com/mcaci/msdb5/v2/app/player"
)

type Opts struct {
	Players player.Players
}

func Cond(g *Opts) bool {
	// no more cards to play
	return g.Players.All(player.EmptyHanded)
}
