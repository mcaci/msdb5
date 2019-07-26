package phase

import (
	"fmt"

	"github.com/mcaci/msdb5/dom/player"
)

type companioner interface {
	Find(player.Predicate) (int, *player.Player)
}

func CardAction(rq cardProvider, comp companioner) Data {
	c, err := rq.Card()
	idx, _ := comp.Find(player.IsCardInHand(c))
	var errCardNotInHand error
	if idx < 0 {
		errCardNotInHand = fmt.Errorf("Card is not in players hand")
	}
	return Data{card: c, plIdx: uint8(idx), cardNotFound: err, cardNotInHand: errCardNotInHand}
}
