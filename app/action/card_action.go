package action

import (
	"errors"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

var errCardNotInHand = errors.New("Card not in hand")

type cardValueProvider interface {
	Card() (*card.Item, error)
	Value() string
}

type data struct {
	card *card.Item
	pl   *player.Player
}

type actioner interface {
	act(data)
	notAcceptedZeroErr() error
	pls() team.Players
}

func cardAction(rq cardValueProvider, a actioner) error {
	if rq.Value() == "0" {
		return a.notAcceptedZeroErr()
	}
	c, err := rq.Card()
	idx, pl := a.pls().Find(player.IsCardInHand(*c))
	if err == nil && idx < 0 {
		return errCardNotInHand
	}
	a.act(data{c, pl})
	return nil
}
