package team

import (
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// BriscolaTeam struct
type BriscolaTeam []player.Scorer

// Add func
func (team *BriscolaTeam) Add(players ...player.Scorer) {
	*team = append(*team, players...)
}

// Score func
func (team BriscolaTeam) Score() (total uint8) {
	for _, player := range team {
		total += player.Count(briscola.Points)
	}
	return
}
