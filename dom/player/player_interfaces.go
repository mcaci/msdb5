package player

import "github.com/nikiforosFreespirit/msdb5/dom/card"

// Predicate type
type Predicate func(p *Player) bool

// Scorer interface
type Scorer interface {
	Count(scorer func(card.ID) uint8) uint8
}

// Counter interface
type Counter interface {
	FoldedChecker
	EmptyHandChecker
	EmptyNameChecker
}

// FoldedChecker interface
type FoldedChecker interface {
	Folded() bool
}

// EmptyNameChecker interface
type EmptyNameChecker interface {
	IsNameEmpty() bool
}

// EmptyHandChecker interface
type EmptyHandChecker interface {
	IsHandEmpty() bool
}
