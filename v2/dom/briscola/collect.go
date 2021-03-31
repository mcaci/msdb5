package briscola

import (
	"github.com/mcaci/ita-cards/set"
)

func Collect(from, to interface{ Pile() *set.Cards }) {
	set.Move(from.Pile(), to.Pile())
}
