package team

import (
	"errors"
	"fmt"
	"log"

	"github.com/mcaci/msdb5/v2/dom/player"
)

// Players is a slice of Players
type Players []player.Player

// Registrator registers Players with their names
type Registrator func(string) error

// New creates new container for players
func New(nPlayers int) *Players {
	players := make(Players, nPlayers)
	for i := range players {
		players[i] = player.New(&player.Options{})
	}
	return &players
}

// NewWithRegistrator creates new container for players
func NewWithRegistrator(nPlayers int) (*Players, Registrator) {
	pls := New(nPlayers)
	var i int
	f := func(n string) error {
		if i >= nPlayers {
			return errors.New("noop: max players reached")
		}
		log.Printf("registering player %d with name %q", i, n)
		o := player.Options{Name: n}
		switch nPlayers {
		case 2:
			o.For2P = true
		case 5:
			o.For5P = true
		default:
			return fmt.Errorf("%d players not supported", nPlayers)
		}
		(*pls)[i] = player.New(&o)
		i++
		return nil
	}
	return pls, f
}

func (players *Players) Add(p player.Player) {
	*players = append(*players, p)
}

func (players Players) SelectIndex(prd player.Predicate) (uint8, error) {
	for i, p := range players {
		if !prd(p) {
			continue
		}
		return uint8(i), nil
	}
	return 0, errors.New("not found")
}

// Part partition players in two groups according to a predicate
func (players Players) Part(predicate player.Predicate) (t1, t2 Players) {
	for _, p := range players {
		if predicate(p) {
			t1.Add(p)
			continue
		}
		t2.Add(p)
	}
	return
}

// Count counts the number of players satisfying the predicate
func Count(players Players, predicate player.Predicate) (count uint8) {
	for _, p := range players {
		if predicate(p) {
			count++
		}
	}
	return
}
