package action

import (
	"errors"

	"github.com/mcaci/ita-cards/set"
)

type PlayCard struct {
	PlCards *set.Cards
}

func (c PlayCard) exec(plCProv playerCardProvider) {
	cards := plCProv.Pl().Hand()
	index := cards.Find(*plCProv.Card())
	c.PlCards.Add((*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}

func (c PlayCard) notAcceptedZeroErr() error {
	return errors.New("Value 0 for card allowed only for ExchangingCard phase")
}
