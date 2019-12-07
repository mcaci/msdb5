package player

import (
	"github.com/mcaci/ita-cards/card"
)

// Predicate type
type Predicate func(p *Player) bool

// IsCardInHand func
func IsCardInHand(c card.Item) Predicate { return func(p *Player) bool { return p.hand.Find(c) != -1 } }

// MatchingHost func
func MatchingHost(host string) Predicate { return func(p *Player) bool { return p.host == host } }

// Matching func
func Matching(o *Player) Predicate { return func(p *Player) bool { return p == o } }

// IsNameEmpty var
var IsNameEmpty Predicate = func(p *Player) bool { return p.name == "" }

// IsHandEmpty var
var IsHandEmpty Predicate = func(p *Player) bool { return len(p.hand) == 0 }

// Folded var
var Folded Predicate = func(p *Player) bool { return p.fold }
