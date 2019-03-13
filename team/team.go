package team

import (
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// BriscolaTeam struct
type BriscolaTeam struct {
	players []player.ScoreCounter
}

// Add func
func (t *BriscolaTeam) Add(players ...player.ScoreCounter) {
	t.players = append(t.players, players...)
}

// Score func
func (t BriscolaTeam) Score() (total uint8) {
	for _, player := range t.players {
		total += player.Count(briscola.Points)
	}
	return
}
