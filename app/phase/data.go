package phase

import (
	"github.com/mcaci/msdb5/dom/auction"
)

type valueProvider interface{ Value() string }

type Data struct {
	name string

	toFold bool
	score  auction.Score
}

func (d Data) Name() string { return d.name }

func (d Data) ToFold() bool         { return d.toFold }
func (d Data) Score() auction.Score { return d.score }
