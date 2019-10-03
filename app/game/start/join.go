package start

import (
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

// Join func
func Join(g interface{ Players() team.Players }, origin string, channel chan []byte) {
	if i, p := g.Players().Find(player.MatchingHost("")); i != -1 {
		p.Join(origin)
		p.Attach(channel)
	}
}
