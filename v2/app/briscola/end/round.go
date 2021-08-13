package end

import (
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola/player"
)

type Opts struct {
	Players briscola.Players
}

func Cond(g *Opts) bool {
	// no more cards to play
	return g.Players.All(player.EmptyHanded)
}
