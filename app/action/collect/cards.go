package collect

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
)

type Info struct {
	current   *player.Player
	toCollect *set.Cards
}

func NewInfo(current *player.Player, toCollect *set.Cards) *Info {
	return &Info{current, toCollect}
}

func (c Info) CurrentPlayer() *player.Player { return c.current }
func (c Info) Cards() *set.Cards             { return c.toCollect }

type infoProvider interface {
	CurrentPlayer() *player.Player
	Cards() *set.Cards
}

func Played(g infoProvider) {
	lastPlayerPile := g.CurrentPlayer().Pile()
	if len(*g.Cards()) == 5 {
		set.Move(g.Cards(), lastPlayerPile)
	}
}
