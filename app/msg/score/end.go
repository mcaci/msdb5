package score

import (
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/team"
)

// Calc func
func Calc(g team.Callers, players team.Players) (totalTeam1, totalTeam2 uint8) {
	for _, p := range players {
		score := p.Pile().Sum(briscola.Points)
		if team.IsInCallers(g, p) {
			totalTeam1 += score
			continue
		}
		totalTeam2 += score
	}
	return
}
