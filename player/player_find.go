package player

import "errors"

// Find func
func Find(nameOrHost string, players []*Player) (*Player, error) {
	isInfoPresent := func(p *Player) bool { return p.Name() == nameOrHost || p.Host() == nameOrHost }
	for _, p := range players {
		if isInfoPresent(p) {
			return p, nil
		}
	}
	return nil, errors.New("Player " + nameOrHost + " not found")
}

