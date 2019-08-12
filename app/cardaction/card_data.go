package cardaction

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

type CardData struct {
	card    *card.Item
	pl      *player.Player
	cardErr error
}

func (d CardData) Card() *card.Item   { return d.card }
func (d CardData) Pl() *player.Player { return d.pl }
func (d CardData) CardErr() error     { return d.cardErr }
