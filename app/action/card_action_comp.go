package action

import (
	"errors"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

type Comp struct {
	SetC func(*card.Item)
	SetP func(*player.Player)
}

func (c Comp) exec(plCProv playerCardProvider) {
	c.SetC(plCProv.Card())
	c.SetP(plCProv.Pl())
}

func (c Comp) notAcceptedZeroErr() error {
	return errors.New("Value 0 for card allowed only for ExchangingCard phase")
}
