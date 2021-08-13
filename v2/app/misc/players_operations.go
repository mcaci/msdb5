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
		if !predicate(p) {
			continue
		}
		return uint8(i), nil
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

// Part partition players in two groups according to a predicate
func (players Players) Part(predicate Predicate) (t1, t2 Players) {
	for _, p := range players {
		if predicate(p) {
			t1.Add(p)
			continue
		}
		t2.Add(p)
	}
	return
}

// Count counts the number of players satisfying the predicate
func Count(players Players, predicate Predicate) (count uint8) {
	for _, p := range players {
		if predicate(p) {
			count++
		}
	}
	return
}
