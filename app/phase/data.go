package phase

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
)

type valueProvider interface{ Value() string }

type cardProvider interface{ Card() (*card.Item, error) }

type Data struct {
	name string

	toFold bool
	score  auction.Score
}

type CardData struct {
	card    *card.Item
	pl      *player.Player
	cardErr error
}

func (d Data) Name() string { return d.name }

func (d Data) ToFold() bool         { return d.toFold }
func (d Data) Score() auction.Score { return d.score }

func (d CardData) Card() *card.Item   { return d.card }
func (d CardData) Pl() *player.Player { return d.pl }
func (d CardData) CardErr() error     { return d.cardErr }
