package team

import (
	"errors"

	"github.com/mcaci/msdb5/dom/player"
)

type expectedPlayerInterface interface {
	CurrentPlayer() *player.Player
	SenderInformation
}

func CheckOrigin(g expectedPlayerInterface) error {
	var ErrUnexpectedPlayer = errors.New("Unexpected player")
	matchHost := player.MatchingHost(g.From())
	matchPl := player.Matching(g.CurrentPlayer())
	criteria := func(p *player.Player) bool { return matchPl(p) && matchHost(p) }
	if i, _ := g.Players().Find(criteria); i == -1 {
		return ErrUnexpectedPlayer
	}
	return nil
}
