package team

import (
	"github.com/mcaci/msdb5/dom/player"
)

// Players struct
type Players []*player.Player

// Add func
func (playerSet *Players) Add(p *player.Player) {
	*playerSet = append(*playerSet, p)
}

// Find func
func (playerSet Players) Find(predicate player.Predicate) (int, *player.Player) {
	for i, p := range playerSet {
		if predicate(p) {
			return i, p
		}
	}
	return -1, nil
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
	for _, p := range playerSet {
		if predicate(p) {
			return false
		}
	}
	return true
}

// Part func
func (playerSet Players) Part(pred player.Predicate) (t1, t2 Players) {
	for _, p := range playerSet {
		if pred(p) {
			t1.Add(p)
			continue
		}
		t2.Add(p)
	}
	return
}
