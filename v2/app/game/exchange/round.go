package exchange

import "github.com/mcaci/ita-cards/set"

func Round(g struct {
	Hand, Side *set.Cards
	hIdx, sIdx int
}) {
	(*g.Hand)[g.hIdx], (*g.Side)[g.sIdx] = (*g.Side)[g.sIdx], (*g.Hand)[g.hIdx]
}
