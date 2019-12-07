package team

import (
	"errors"

	"github.com/mcaci/msdb5/dom/player"
)

// ErrPlayerNotFound error
var ErrPlayerNotFound = errors.New("Player not found")

// MustFind func
func (playerSet Players) MustFind(predicate player.Predicate) uint8 {
	i, err := playerSet.Index(predicate)
	if err != nil {
		panic(err)
	}
	return i
}

// All func
func (playerSet Players) All(predicate player.Predicate) bool {
	for _, p := range playerSet {
		if !predicate(p) {
			return false
		}
	}
	return true
}

// None func
func (playerSet Players) None(predicate player.Predicate) bool {
	_, err := playerSet.Index(predicate)
	return err != nil
}

// CheckOrigin func
func CheckOrigin(players Players, senderHost string, expected *player.Player) error {
	senderMatch := player.MatchingHost(senderHost)
	gamePlayerMatch := player.Matching(expected)
	criteria := func(p *player.Player) bool { return senderMatch(p) && gamePlayerMatch(p) }
	if players.None(criteria) {
		return ErrPlayerNotFound
	}
	return nil
}

// Index func
func (playerSet Players) Index(predicate player.Predicate) (uint8, error) {
	for i, p := range playerSet {
		if predicate(p) {
			return uint8(i), nil
		}
	}
	return 0, ErrPlayerNotFound
}
