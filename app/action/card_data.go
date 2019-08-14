package action

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

type data struct {
	card *card.Item
	pl   *player.Player
}

func (d data) Card() *card.Item   { return d.card }
func (d data) Pl() *player.Player { return d.pl }
