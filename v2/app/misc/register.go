package misc

import (
	"errors"
	"fmt"
	"log"
)

// Registrator registers Players with their names
type Registrator func(string) error

// NewWithRegistrator creates new container for players
func NewWithRegistrator(nPlayers int) (*Players, Registrator) {
	pls := NewPlayers(nPlayers)
	var i int
	f := func(n string) error {
		if i >= nPlayers {
			return errors.New("noop: max players reached")
		}
		log.Printf("registering player %d with name %q", i, n)
		o := Options{Name: n}
		switch nPlayers {
		case 2:
			o.For2P = true
		case 5:
			o.For5P = true
		default:
			return fmt.Errorf("%d players not supported", nPlayers)
		}
		(*pls)[i] = New(&o)
		i++
		return nil
	}
	return pls, f
}
