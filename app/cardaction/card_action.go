package cardaction

import (
	"errors"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
)

// ErrCardNotInHand error
var ErrCardNotInHand = errors.New("Card not in hand")

type cardActioner interface {
	Find(player.Predicate) (int, *player.Player)
}

type Actioner interface {
	exec(plCProv playerCardProvider)
	notAcceptedZeroErr() error
}

type cardValueProvider interface {
	Card() (*card.Item, error)
	Value() string
}

func CardAction(rq cardValueProvider, act cardActioner, a Actioner) error {
	if rq.Value() == "0" {
		return a.notAcceptedZeroErr()
	}
	c, err := rq.Card()
	idx, p := act.Find(player.IsCardInHand(*c))
	if err == nil && idx < 0 {
		err = ErrCardNotInHand
	}
	data := CardData{card: c, pl: p, cardErr: err}
	if err != nil {
		return err
	}
	a.exec(data)
	return nil
}

type playerCardProvider interface {
	Card() *card.Item
	Pl() *player.Player
}

type Comp struct {
	SetC func(*card.Item)
	SetP func(*player.Player)
}

func (c Comp) exec(plCProv playerCardProvider) {
	c.SetC(plCProv.Card())
	c.SetP(plCProv.Pl())
}
func (c Comp) notAcceptedZeroErr() error { return nil }

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
	return errors.New("Value 0 for card allowed only for ExchangingCard phase")
}

type Play struct {
	PlCards *set.Cards
}

func (c Play) exec(plCProv playerCardProvider) {
	cards := plCProv.Pl().Hand()
	index := cards.Find(*plCProv.Card())
	c.PlCards.Add((*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}
func (c Play) notAcceptedZeroErr() error {
	return errors.New("Value 0 for card allowed only for ExchangingCard phase")
}
