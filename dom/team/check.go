package team

import (
	"errors"

	"github.com/mcaci/msdb5/dom/player"
)

type expectedPlayerInterface interface {
	CurrentPlayer() *player.Player
	Players() Players
}

func CheckOrigin(g expectedPlayerInterface, rq interface{ From() string }) error {
	var ErrUnexpectedPlayer = errors.New("Unexpected player")
	matchHost := player.MatchingHost(rq.From())
	matchPl := player.Matching(g.CurrentPlayer())
	criteria := func(p *player.Player) bool { return matchPl(p) && matchHost(p) }
	if i, _ := g.Players().Find(criteria); i == -1 {
		return ErrUnexpectedPlayer
	}
	return nil
}
