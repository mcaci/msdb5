package player

import "github.com/mcaci/ita-cards/card"

// IsCardInHand func
func IsCardInHand(c card.Item) Predicate {
	return func(p *Player) bool { return p.Has(c) }
}

// Has func
func (player *Player) Has(id card.Item) bool { return player.hand.Find(id) != -1 }

// IsSameHost func
func (player Player) IsSameHost(origin string) bool { return player.host == origin }

// Predicate type
type Predicate func(p *Player) bool

// IsNameEmpty var
var IsNameEmpty Predicate = func(p *Player) bool { return p.name == "" }

// IsHandEmpty var
var IsHandEmpty Predicate = func(p *Player) bool { return len(p.hand) == 0 }

// Folded var
var Folded Predicate = func(p *Player) bool { return p.fold }
