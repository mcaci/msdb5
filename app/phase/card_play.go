package phase

import (
	"fmt"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
)

type localCardProvider interface {
	Card() (card.ID, error)
}

// CardAction func
func CardAction(rq localCardProvider, cards, to *deck.Cards, effect func(cards, to *deck.Cards, index, toIndex int)) error {
	c, err := rq.Card()
	if err != nil {
		return err
	}
	index := cards.Find(c)
	if index == -1 {
		return fmt.Errorf("Card is not in players hand")
	}
	effect(cards, to, index, 0)
	return nil
}
