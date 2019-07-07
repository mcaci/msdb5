package play

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

// Exchange func
func Exchange(fromCard card.ID, from, to *deck.Cards) error {
	fromCardIndex := from.Find(fromCard)
	if fromCardIndex == -1 {
		return fmt.Errorf("Card is not in players hand")
	}
	toCardIndex := 0
	toCard := (*to)[toCardIndex]
	from.Add(toCard)
	to.Remove(toCardIndex)
	to.Add(fromCard)
	from.Remove(fromCardIndex)
	return nil
}
