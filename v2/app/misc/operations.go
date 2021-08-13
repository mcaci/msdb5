package misc

import (
	"errors"
)

// ErrPlayerNotFound error
var ErrPlayerNotFound = errors.New("player not found")

// Index returns the index of the player that satisfies the predicate
// or an error if not found
func (playerSet Players) Index(predicate Predicate) (uint8, error) {
	for i, p := range playerSet {
		if predicate(p) {
			return uint8(i), nil
		}
	}
	return 0, ErrPlayerNotFound
}

// All verifies if all players satisfy the predicate
func (playerSet Players) All(predicate Predicate) bool {
	_, err := playerSet.Index(func(p Player) bool { return !predicate(p) })
	return err != nil
}

// None verifies if no player satisfies the predicate
func (playerSet Players) None(predicate Predicate) bool {
	_, err := playerSet.Index(predicate)
	return err != nil
}
