package briscola

import (
	"github.com/mcaci/ita-cards/card"
)

type cardIds []uint8

func (ids cardIds) Len() int { return len(ids) }

func (ids cardIds) Less(i, j int) bool {
	return isOtherHigher(*card.MustID(ids[i]), *card.MustID(ids[j]))
}

func (ids cardIds) Swap(i, j int) { ids[i], ids[j] = ids[j], ids[i] }
