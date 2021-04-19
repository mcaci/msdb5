package session

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/app/briscola5"
)

type Briscola struct {
	Game *briscola.Game
	Deck set.Cards
}

type Briscola5 struct {
	Game *briscola5.Game
}
