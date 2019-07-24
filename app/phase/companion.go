package phase

import (
	"github.com/mcaci/msdb5/dom/player"
)

type companioner interface {
	Find(player.Predicate) (int, *player.Player)
}

func Companion(rq cardProvider, comp companioner) Data {
	c, err := rq.Card()
	// setBriscolaCard(c)
	idx, _ := comp.Find(player.IsCardInHand(c))
	// setCompanion(pl)
	return Data{briscola: c, comp: uint8(idx), cardNotFound: err}
}
