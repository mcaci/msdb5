package collect

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type AllInfo struct {
	infoProvider
	players team.Players
}

func NewAllInfo(current *player.Player, toCollect *set.Cards, players team.Players) *AllInfo {
	return &AllInfo{Info{current, toCollect}, players}
}

func (c AllInfo) Players() team.Players { return c.players }

type allInfoProvider interface {
	infoProvider
	Players() team.Players
}

func All(g allInfoProvider) {
	lastPlayerPile := g.CurrentPlayer().Pile()
	if len(*g.CurrentPlayer().Hand()) > 0 {
		for _, pl := range g.Players() {
			set.Move(pl.Hand(), lastPlayerPile)
		}
	}
	if len(*g.Cards()) > 0 {
		set.Move(g.Cards(), lastPlayerPile)
	}
}
