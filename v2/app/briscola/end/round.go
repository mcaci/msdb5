package end

import "github.com/mcaci/msdb5/v2/app/misc"

type Opts struct {
	Players misc.Players
}

func Cond(g *Opts) bool {
	// no more cards to play
	return g.Players.All(misc.EmptyHanded)
}
