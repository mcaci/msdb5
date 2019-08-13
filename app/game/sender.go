package game

import (
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func senderIndex(g interface{ Players() team.Players }, rq interface{ From() string }) int {
	index, _ := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(rq.From()) })
	return index
}

func sender(g interface{ Players() team.Players }, rq interface{ From() string }) *player.Player {
	return g.Players()[senderIndex(g, rq)]
}
