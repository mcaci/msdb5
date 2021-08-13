package end

type Opts struct {
	Players misc.Players
}

func Cond(g *Opts) bool {
	// no more cards to play
	return g.Players.All(misc.EmptyHanded)
}
