package set

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/player"
)

// Find func
func (players Players) Find(nameOrHost string) (*player.Player, error) {
	isInfoPresent := func(p *player.Player) bool { return p.Name() == nameOrHost || p.Host() == nameOrHost }
	for _, p := range players {
		if isInfoPresent(p) {
			return p, nil
		}
	}
	return nil, errors.New("Player " + nameOrHost + " not found")
}
