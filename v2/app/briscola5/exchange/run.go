package exchange

import (
	"math/rand"

	"github.com/mcaci/ita-cards/set"
)

func Run(g struct {
	Hand, Side *set.Cards
}) {
	for {
		idx := rand.Intn(len(*g.Hand))
		if idx > 2 {
			break
		}
		Round(struct {
			Hand, Side *set.Cards
			hIdx, sIdx int
		}{
			Hand: g.Hand, Side: g.Side,
			hIdx: idx, sIdx: 0,
		})
	}
}
