package misc

import (
	"github.com/mcaci/ita-cards/card"
)


// IsCardInHand func
func IsCardInHand(c card.Item) func(p Player) bool {
	return func(p Player) bool {
		return p.Hand().Find(c) != -1
	}
}

// EmptyHanded func
func EmptyHanded(p Player) bool { return len(*p.Hand()) == 0 }
