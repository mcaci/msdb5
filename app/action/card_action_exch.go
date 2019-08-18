package action

import (
	"github.com/mcaci/ita-cards/set"
)

type Exch struct {
	Side *set.Cards
}

func (c Exch) exec(plCProv playerCardProvider) {
	cards := plCProv.Pl().Hand()
	index := cards.Find(*plCProv.Card())
	toCards := c.Side
	awayCard := (*cards)[index]
	(*cards)[index] = (*toCards)[0]
	*toCards = append((*toCards)[1:], awayCard)
}

func (c Exch) notAcceptedZeroErr() error {
	return nil
}
