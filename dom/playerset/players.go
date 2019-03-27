package playerset

import (
	"errors"
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Players struct
type Players []*player.Player

// Add func
func (playerSet *Players) Add(p player.Player) {
	*playerSet = append(*playerSet, &p)
}

// Find func
func (playerSet Players) Find(predicate func(p *player.Player) bool) (*player.Player, error) {
	for _, p := range playerSet {
		if predicate(p) {
			return p, nil
		}
	}
	return nil, errors.New("Player not found")
}

// FindIndex func
func (playerSet Players) FindIndex(predicate func(p *player.Player) bool) (int, error) {
	for i, p := range playerSet {
		if predicate(p) {
			return i, nil
		}
	}
	return -1, errors.New("Player not found")
}

// Count func
func (playerSet Players) Count(predicate func(p *player.Player) bool) (count uint8) {
	for _, p := range playerSet {
		if predicate(p) {
			count++
		}
	}
	return
}

// All func
func (playerSet Players) All(predicate func(p *player.Player) bool) bool {
	for _, p := range playerSet {
		if !predicate(p) {
			return false
		}
	}
	return true
}

func (playerSet Players) String() (str string) {
	for _, p := range playerSet {
		str += fmt.Sprintf("- %+v -", *p)
	}
	return

}
