package start

import (
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

// Join func
func Join(g interface{ Players() team.Players }, origin string, channel chan []byte) {
	if i, err := g.Players().Index(player.MatchingHost("")); err == nil {
		p := g.Players().At(i)
		p.Join(origin)
		p.Attach(channel)
	}
}
