package player

import (
	"errors"
)

// Find func
func (players Players) Find(nameOrHost string) (*Player, error) {
	isInfoPresent := func(p *Player) bool { return p.Name() == nameOrHost || p.Host() == nameOrHost }
	for _, p := range players {
		if isInfoPresent(p) {
			return p, nil
		}
	}
	return nil, errors.New("Player " + nameOrHost + " not found")
}
