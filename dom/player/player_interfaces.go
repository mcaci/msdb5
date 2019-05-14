package player

import "github.com/nikiforosFreespirit/msdb5/dom/card"

// Scorer interface
type Scorer interface {
	Count(scorer func(card.ID) uint8) uint8
}
