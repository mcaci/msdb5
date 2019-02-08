package playerset

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/player"
)

// Players struct
type Players []*player.Player

// Add func
func (set *Players) Add(p player.Player) {
	*set = append(*set, &p)
}

// Find func
func (set Players) Find(nameOrHost string) (*player.Player, error) {
	isInfoPresent := func(p *player.Player) bool { return p.Name() == nameOrHost || p.Host() == nameOrHost }
	for _, p := range set {
		if isInfoPresent(p) {
			return p, nil
		}
	}
	return nil, errors.New("Player " + nameOrHost + " not found")
}
func (set Players) String() string {
	var str string
	for _, player := range set {
		str += player.String() + " "
	}
	return str
}
