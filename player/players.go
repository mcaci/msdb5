package player

import "errors"

// Players struct
type Players []*Player

// Add func
func (set *Players) Add(p Player) {
	*set = append(*set, &p)
}

func (set Players) String() string {
	var str string
	for _, player := range set {
		str += player.String() + " "
	}
	return str
}

// Find func
func (set Players) Find(nameOrHost string) (*Player, error) {
	isInfoPresent := func(p *Player) bool { return p.Name() == nameOrHost || p.Host() == nameOrHost }
	for _, p := range set {
		if isInfoPresent(p) {
			return p, nil
		}
	}
	return nil, errors.New("Player " + nameOrHost + " not found")
}
