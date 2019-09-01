package action

import (
	"github.com/mcaci/msdb5/dom/player"
)

type joinData struct {
	currentPlayer *player.Player
}

func (j joinData) valueSet(val string) {
	j.currentPlayer.RegisterAs(val)
}
