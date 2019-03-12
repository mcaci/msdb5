package player

import "github.com/nikiforosFreespirit/msdb5/card"

// ScoreCounter interface
type ScoreCounter interface {
	Count(scorer func(card.ID) uint8) uint8
}
