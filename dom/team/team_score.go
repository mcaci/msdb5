package team

import (
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/deck"
)

// Piler interface
type Piler interface {
	Pile() *deck.Cards
}

// Score func
func Score(caller, companion Piler, players []Piler) (totalTeam1, totalTeam2 uint8) {
	for _, pl := range players {
		score := pl.Pile().Sum(briscola.Points)
		if pl == caller || pl == companion {
			totalTeam1 += score
			continue
		}
		totalTeam2 += score
	}
	return
}
