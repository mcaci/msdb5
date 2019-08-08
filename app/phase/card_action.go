package phase

import (
	"errors"

	"github.com/mcaci/msdb5/dom/player"
)

// ErrCardNotInHand error
var ErrCardNotInHand = errors.New("Card not in hand")

type cardActioner interface {
	Find(player.Predicate) (int, *player.Player)
}

func CardAction(rq cardProvider, act cardActioner) CardData {
	c, err := rq.Card()
	idx, p := act.Find(player.IsCardInHand(c))
	if err == nil && idx < 0 {
		err = ErrCardNotInHand
	}
	return CardData{card: c, plIdx: uint8(idx), pl: p, cardErr: err}
}
