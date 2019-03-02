package playerset

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/player"
)

// Players struct
type Players []*player.Player

// Add func
func (playerSet *Players) Add(p player.Player) {
	*playerSet = append(*playerSet, &p)
}

// Find func
func (playerSet Players) Find(isInfoPresent func(p *player.Player) bool) (*player.Player, error) {
	for _, p := range playerSet {
		if isInfoPresent(p) {
			return p, nil
		}
	}
	return nil, errors.New("Player not found")
}

func (playerSet Players) String() (str string) {
	for _, p := range playerSet {
		str += p.String() + " "
	}
	return
}
