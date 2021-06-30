package player

import (
	"github.com/mcaci/ita-cards/card"
)

// Predicate type
type Predicate func(p Player) bool

// IsCardInHand func
func IsCardInHand(c card.Item) Predicate {
	return func(p Player) bool {
		return p.Hand().Find(c) != -1
	}
}

// EmptyHanded var
var EmptyHanded Predicate = func(p Player) bool { return len(*p.Hand()) == 0 }

// IsInCallers func
func IsInCallers(t interface {
	Caller() Player
	Companion() Player
}) Predicate {
	return func(p Player) bool { return eq(p, t.Caller()) || eq(p, t.Companion()) }
}

func eq(p, q Player) bool {
	return p.Name() == q.Name() && p.Hand() == q.Hand() && p.Pile() == q.Pile()
}

// Folded var
var Folded Predicate = func(p Player) bool {
	b5p, ok := p.(*B5Player)
	if !ok {
		return false
	}
	return b5p.fold
}

// NotFolded var
var NotFolded = func(p Player) bool {
	b5p, ok := p.(*B5Player)
	if !ok {
		return false

	}
	return !b5p.fold
}
