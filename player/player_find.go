package player

import "errors"

// Find func
func Find(nameOrHost string, players []*Player) (player *Player, err error) {
	player, err = ByName(nameOrHost, players)
	if err != nil {
		player, err = ByHost(nameOrHost, players)
	}
	if err != nil {
		err = errors.New("Player " + nameOrHost + " not found")
	}
	return
}

func ByName(name string, players []*Player) (player *Player, err error) {
	var found bool
	for _, p := range players {
		found = p.Name() == name
		if found {
			player = p
			break
		}
	}
	if !found {
		err = errors.New("Player " + name + " not found")
	}
	return
}

func ByHost(host string, players []*Player) (player *Player, err error) {
	var found bool
	for _, p := range players {
		found = p.Host() == host
		if found {
			player = p
			break
		}
	}
	if !found {
		err = errors.New("Player " + host + " not found")
	}
	return
}
