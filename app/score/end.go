package score

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

// Piler interface
type Piler interface {
	Pile() *set.Cards
}

// Calc func
func Calc(caller, companion Piler, players []Piler, cardValuer func(id card.Item) uint8) (totalTeam1, totalTeam2 uint8) {
	for _, pl := range players {
		score := pl.Pile().Sum(cardValuer)
		if pl == caller || pl == companion {
			totalTeam1 += score
			continue
		}
		totalTeam2 += score
	}
	return
}
