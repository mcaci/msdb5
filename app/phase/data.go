package phase

import (
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
)

type valueProvider interface{ Value() string }

type cardProvider interface{ Card() (card.ID, error) }

type Data struct {
	name string

	toFold    bool
	score     auction.Score
	sideCards uint8

	briscola     card.ID
	comp         uint8
	cardNotFound error
}

func (d Data) Name() string { return d.name }

func (d Data) ToFold() bool         { return d.toFold }
func (d Data) Score() auction.Score { return d.score }
func (d Data) SideCards() uint8     { return d.sideCards }

func (d Data) Briscola() card.ID { return d.briscola }
func (d Data) CompIdx() uint8    { return d.comp }
