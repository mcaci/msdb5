package action

import (
	"errors"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

// ErrCardNotInHand error
var ErrCardNotInHand = errors.New("Card not in hand")

type cardValueProvider interface {
	Card() (*card.Item, error)
	Value() string
}

type actioner interface {
	exec(plCProv playerCardProvider)
	notAcceptedZeroErr() error
}

type playerCardProvider interface {
	Card() *card.Item
	Pl() *player.Player
}

func CardAction(rq cardValueProvider, finder interface {
	Find(player.Predicate) (int, *player.Player)
}, a actioner) error {
	if rq.Value() == "0" {
		return a.notAcceptedZeroErr()
	}
	c, err := rq.Card()
	idx, p := finder.Find(player.IsCardInHand(*c))
	if err == nil && idx < 0 {
		err = ErrCardNotInHand
	}
	d := data{card: c, pl: p}
	if err != nil {
		return err
	}
	a.exec(d)
	return nil
}
