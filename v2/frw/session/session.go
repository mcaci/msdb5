package session

import (
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
}

func NewBriscola() *Briscola {
	return &Briscola{
		Game:  briscola.New(),
		Ready: make(chan interface{}),
	}
}

type Briscola5 struct {
	Game *briscola5.Game
}
