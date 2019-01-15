package board

import (
	"log"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player/set"
)

// Players func
func (b *Board) Players() set.Players {
	return b.players
}

// PChans func
func (b *Board) PChans() []chan card.ID {
	return b.pChans
}

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
