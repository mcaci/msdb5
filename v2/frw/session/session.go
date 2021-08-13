package session

import (
	"sync"

	"github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/app/briscola5"
)

const (
	NPlBriscola = 2
	// nPlBriscola5 = 5
)

type Briscola struct {
	Game  *briscola.Game
	Curr  uint8
	NPls  uint8
	Ready chan interface{}
	Wg    *sync.WaitGroup
}

func NewBriscola() *Briscola {
	b := &Briscola{
		Game:  briscola.NewGame(briscola.WithDefaultOptions),
		Ready: make(chan interface{}),
		Wg:    &sync.WaitGroup{},
	}
	b.Wg.Add(2)
	return b
}

func (b *Briscola) GetAndIncr() int {
	npls := b.NPls
	b.NPls++
	return int(npls)
}

type Briscola5 struct {
	Game *briscola5.Game
}
