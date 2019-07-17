package team

import (
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/card"
)

// Scorer interface
type Scorer interface {
	Points(scorer func(card.ID) uint8) uint8
}

// Score func
func Score(caller, companion Scorer, players ...Scorer) (totalTeam1, totalTeam2 uint8) {
	for _, player := range players {
		score := player.Points(briscola.Points)
		if player == caller || player == companion {
			totalTeam1 += score
			continue
		}
		totalTeam2 += score
	}
	return
}
