package score

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/team"
)

// Calc func
func Calc(g team.Callers, players team.Players, cardValuer func(id card.Item) uint8) (totalTeam1, totalTeam2 uint8) {
	for _, p := range players {
		score := p.Pile().Sum(cardValuer)
		if team.IsInCallers(g, p) {
			totalTeam1 += score
			continue
		}
		totalTeam2 += score
	}
	return
}
