package briscola

import (
	"errors"
	"log"

	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type Players struct{ team.Players }

// NewPlayers creates new container for briscola5 players
func NewPlayers(nPlayers int) *Players {
	players := make(team.Players, nPlayers)
	for i := range players {
		players[i] = player.New(&player.Options{For2P: true}).(*player.Player)
	}
	return &Players{Players: players}
}

func (pls *Players) Registration() func(string) error {
	const nPlayers = 2
	var i int
	return func(s string) error {
		if i >= nPlayers {
			return errors.New("noop: max players reached")
		}
		log.Printf("registering player %d with name %q", i, s)
		pls.At(i).RegisterAs(s)
		i++
		return nil
	}
}

func (pls *Players) List() team.Players      { return pls.Players }
func (pls *Players) Len() int                { return len(pls.Players) }
func (pls *Players) At(i int) *player.Player { return pls.Players[i] }

func (pls *Players) Select(prd player.Predicate) (*player.Player, error) {
	i, err := pls.SelectIndex(prd)
	return pls.Players[i], err
}

func (pls *Players) SelectIndex(prd player.Predicate) (uint8, error) {
	for i, p := range pls.Players {
		if !prd(p) {
			continue
		}
		return uint8(i), nil
	}
	return 0, errors.New("not found")
}
