package game

import "github.com/mcaci/msdb5/dom/player"

// Join func
func (g *Game) Join(origin string, channel chan []byte) {
	if i, p := g.players.Find(player.MatchingHost("")); i != -1 {
		p.Join(origin)
		p.Attach(channel)
	}
}
