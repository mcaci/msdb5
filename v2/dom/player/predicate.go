package player

import (
	"github.com/mcaci/ita-cards/card"
)

// Predicate type
type Predicate func(p *Player) bool

// IsCardInHand func
func IsCardInHand(c card.Item) Predicate { return func(p *Player) bool { return p.hand.Find(c) != -1 } }

// Matching func
func Matching(o *Player) Predicate { return func(p *Player) bool { return p == o } }

// EmptyHanded var
var EmptyHanded Predicate = func(p *Player) bool { return len(p.hand) == 0 }
