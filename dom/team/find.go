package team

import (
	"errors"

	"github.com/mcaci/msdb5/dom/player"
)

// ErrPlayerNotFound error
var ErrPlayerNotFound = errors.New("Player not found")

// Index func
func (playerSet Players) Index(predicate player.Predicate) (uint8, error) {
	for i, p := range playerSet {
		if predicate(p) {
			return uint8(i), nil
		}
	}
	return 0, ErrPlayerNotFound
}

// MustIndex func
func (playerSet Players) MustIndex(predicate player.Predicate) uint8 {
	i, err := playerSet.Index(predicate)
	if err != nil {
		panic(err)
	}
	return i
}

// All func
func (playerSet Players) All(predicate player.Predicate) bool {
	_, err := playerSet.Index(func(p *player.Player) bool { return !predicate(p) })
	return err != nil
}

// None func
func (playerSet Players) None(predicate player.Predicate) bool {
	_, err := playerSet.Index(predicate)
	return err != nil
}
