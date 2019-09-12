package game

import "github.com/mcaci/msdb5/dom/player"

// Join func
func (g *Game) Join(origin string, channel chan []byte) {
	for _, p := range g.players {
		if player.IsHostEmpty(p) {
			p.Join(origin)
			p.Attach(channel)
			break
		}
	}
}
