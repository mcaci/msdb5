package briscola

import (
	"errors"
)

// ErrPlayerNotFound error
var ErrPlayerNotFound = errors.New("player not found")

// Predicate type
type Predicate func(p Player) bool

// Index returns the index of the player that satisfies the predicate
// or an error if not found
func (playerSet Players) Index(predicate Predicate) (uint8, error) {
	for i, p := range playerSet {
		if !predicate(*p) {
			continue
		}
		return uint8(i), nil
	}
	return 0, ErrPlayerNotFound
}

// None verifies if no player satisfies the predicate
func (playerSet Players) None(predicate Predicate) bool {
	_, err := playerSet.Index(predicate)
	return err != nil
}
