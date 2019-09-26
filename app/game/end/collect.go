package end

import (
	"github.com/mcaci/ita-cards/set"
)

type collector interface {
	Pile() *set.Cards
}

func Collect(dest collector, orig ...*set.Cards) {
	for _, cards := range orig {
		set.Move(cards, dest.Pile())
	}
}
