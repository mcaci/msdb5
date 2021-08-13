package misc

import (
	"log"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

// Predicate type
type Predicate func(p Player) bool

// IsCardInHand func
func IsCardInHand(c card.Item) func(p Player) bool {
	return func(p Player) bool {
		return p.Hand().Find(c) != -1
	}
}

// EmptyHanded func
func EmptyHanded(p Player) bool { return len(*p.Hand()) == 0 }

// IsInCallers func
func IsInCallers(t interface {
	Caller() Player
	Companion() Player
}) Predicate {
	log.Println(t)
	return func(p Player) bool { return eq(p, t.Caller()) || eq(p, t.Companion()) }
}

func eq(p, q Player) bool {
	return p.Name() == q.Name() && p.Hand() == q.Hand() && p.Pile() == q.Pile()
}

// Folded var
var Folded Predicate = func(p Player) bool {
	b5p, ok := p.(*briscola5.Player)
	if !ok {
		return false
	}
	return b5p.Folded()
}

// NotFolded var
var NotFolded = func(p Player) bool {
	b5p, ok := p.(*briscola5.Player)
	if !ok {
		return false

	}
	return !b5p.Folded()
}
