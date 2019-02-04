package board

import "log"

// Join func
func (b *Board) Join(name, remoteAddr string) {
	for _, player := range b.Players() {
		if player.Name() == "" {
			player.SetName(name)
			player.MyHostIs(remoteAddr)
			return
		}
	}
	log.Println("All players have joined, no further players are expected")
}
