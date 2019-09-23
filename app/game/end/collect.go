package end

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
)

type CollectInfo struct {
	current   *player.Player
	toCollect *set.Cards
}

func NewCollectInfo(current *player.Player, toCollect *set.Cards) *CollectInfo {
	return &CollectInfo{current, toCollect}
}

func (c CollectInfo) CurrentPlayer() *player.Player { return c.current }
func (c CollectInfo) Cards() *set.Cards             { return c.toCollect }

func Collect(g currentPlayerCardsProvider) {
	set.Move(g.Cards(), g.CurrentPlayer().Pile())
}
