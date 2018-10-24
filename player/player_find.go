package player

import "errors"

// Find func
func Find(nameOrHost string, players []*Player) (p *Player, err error) {
	p, found := findInfoIn(players, func(pl *Player) bool { return pl.Name() == nameOrHost })
	if !found {
		p, found = findInfoIn(players, func(pl *Player) bool { return pl.Host() == nameOrHost })
	}
	if !found {
		err = errors.New("Player " + nameOrHost + " not found")
	}
	return
}

func findInfoIn(players []*Player, isInfoPresent func(*Player) bool) (player *Player, found bool) {
	for _, p := range players {
		if isInfoPresent(p) {
			return p, true
		}
	}
	return
}
