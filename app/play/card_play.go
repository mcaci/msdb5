package play

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Play function
func Play(pl *player.Player, card card.ID) error {
	fromCardIndex := pl.Hand().Find(card)
	if fromCardIndex == -1 {
		return fmt.Errorf("Card is not in players hand")
	}
	pl.Hand().Remove(fromCardIndex)
	return nil
}
