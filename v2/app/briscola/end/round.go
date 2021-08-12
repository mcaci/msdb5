package end

import (
	"github.com/mcaci/msdb5/v2/dom/briscola/player"
	"github.com/mcaci/msdb5/v2/dom/briscola/team"
)

type Opts struct {
	Players team.Players
}

func Cond(g *Opts) bool {
	// no more cards to play
	return g.Players.All(player.EmptyHanded)
}
