package exchange

import "github.com/mcaci/ita-cards/set"

func Round(g struct {
	Hand, Side *set.Cards
	hIdx, sIdx int
}) {
	discardedCard := (*g.Hand)[g.hIdx]
	(*g.Hand)[g.hIdx] = (*g.Side)[g.sIdx]
	(*g.Side) = append((*g.Side)[:g.sIdx], (*g.Side)[g.sIdx+1:]...)
	(*g.Side) = append(*g.Side, discardedCard)
}
