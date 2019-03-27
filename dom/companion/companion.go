package companion

import (
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Companion struct
type Companion struct {
	card card.ID
	ref  *player.Player
}

// New func
func New(card card.ID, ref *player.Player) *Companion {
	return &Companion{card, ref}
}

// Card func
func (c *Companion) Card() card.ID {
	return c.card
}

// Ref func
func (c *Companion) Ref() *player.Player {
	return c.ref
}
