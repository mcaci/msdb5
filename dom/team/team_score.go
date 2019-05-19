package team

import (
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Score func
func Score(caller, companion player.Scorer, players ...player.Scorer) (totalTeam1, totalTeam2 uint8) {
	for _, player := range players {
		score := player.Count(briscola.Points)
		if player == caller || player == companion {
			totalTeam1 += score
		} else {
			totalTeam2 += score
		}
	}
	return
}
