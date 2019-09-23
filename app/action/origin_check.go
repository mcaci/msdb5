package action

import (
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type expectedPlayerInterface interface {
	CurrentPlayer() *player.Player
	Players() team.Players

	From() string
}

func CheckOrigin(g expectedPlayerInterface) error {
	matchHost := player.MatchingHost(g.From())
	matchPl := player.Matching(g.CurrentPlayer())
	criteria := func(p *player.Player) bool { return matchPl(p) && matchHost(p) }
	if g.Players().None(criteria) {
		return player.ErrUnexpectedPlayer
	}
	return nil
}
