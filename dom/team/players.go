package team

import (
	"errors"

	"github.com/mcaci/msdb5/dom/player"
)

// Players struct
type Players []*player.Player

// Add func
func (playerSet *Players) Add(p *player.Player) {
	*playerSet = append(*playerSet, p)
}

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

// At func
func (playerSet Players) At(i uint8) *player.Player { return playerSet[i] }

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

// Part func
func (playerSet Players) Part(predicate player.Predicate) (t1, t2 Players) {
	for _, p := range playerSet {
		if predicate(p) {
			t1.Add(p)
			continue
		}
		t2.Add(p)
	}
	return
}
