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

type ExpectedSender struct {
	players team.Players
	origin  string
	p       *player.Player
}

func NewExpectedSender(players team.Players, origin string, p *player.Player) ExpectedSender {
	return ExpectedSender{players, origin, p}
}

func (s ExpectedSender) CurrentPlayer() *player.Player { return s.p }
func (s ExpectedSender) From() string                  { return s.origin }
func (s ExpectedSender) Players() team.Players         { return s.players }

func CheckOrigin(g expectedPlayerInterface) error {
	matchHost := player.MatchingHost(g.From())
	matchPl := player.Matching(g.CurrentPlayer())
	criteria := func(p *player.Player) bool { return matchPl(p) && matchHost(p) }
	if g.Players().None(criteria) {
		return player.ErrUnexpectedPlayer
	}
	return nil
}
