package player

import "errors"

// ByName func
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
