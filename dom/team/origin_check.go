package team

import (
	"github.com/mcaci/msdb5/dom/player"
)

// CheckOrigin func
func CheckOrigin(players Players, senderHost string, expected *player.Player) error {
	senderMatch := player.MatchingHost(senderHost)
	gamePlayerMatch := player.Matching(expected)
	criteria := func(p *player.Player) bool { return senderMatch(p) && gamePlayerMatch(p) }
	if players.None(criteria) {
		return player.ErrUnexpectedPlayer
	}
	return nil
}
