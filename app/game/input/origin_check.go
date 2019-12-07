package input

import (
	"errors"

	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

// ErrUnexpectedPlayerInput error
var ErrUnexpectedPlayerInput = errors.New("Unexpected player playing")

// CheckOrigin func
func CheckOrigin(players team.Players, senderHost string, expected *player.Player) error {
	senderMatch := player.MatchingHost(senderHost)
	gamePlayerMatch := player.Matching(expected)
	criteria := func(p *player.Player) bool { return senderMatch(p) && gamePlayerMatch(p) }
	if players.None(criteria) {
		return ErrUnexpectedPlayerInput
	}
	return nil
}
