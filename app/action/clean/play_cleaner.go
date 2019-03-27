package clean

import (
	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

type CleanUp struct {
	playedCards *deck.Cards
}

func NewCleaner(playedCards *deck.Cards) action.Cleaner {
	return &CleanUp{playedCards}
}

func (c *CleanUp) Clean() {
	c.playedCards.Clear()
}
