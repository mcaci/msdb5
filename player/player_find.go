package player

import "errors"

// Find func
func Find(nameOrHost string, players []*Player) (p *Player, err error) {
	findCriteria := func(p *Player) bool { return p.Name() == nameOrHost || p.Host() == nameOrHost }
	p = findInfoIn(players, findCriteria)
	if p == nil {
		err = errors.New("Player " + nameOrHost + " not found")
	}
	return
}

func findInfoIn(players []*Player, isInfoPresent func(*Player) bool) *Player {
	for _, p := range players {
		if isInfoPresent(p) {
			return p
		}
	}
	return nil
}
