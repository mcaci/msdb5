package briscola

import (
	"errors"
	"log"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type Players struct{ team.Players }

// NewPlayers creates new container for briscola5 players
func NewPlayers() *Players {
	players := make(team.Players, 2)
	for i := range players {
		players[i] = player.New()
	}
	return &Players{Players: players}
}

func (pls *Players) Registration() func(string) error {
	var i int
	return func(s string) error {
		if i >= 5 {
			return errors.New("noop: max players reached")
		}
		log.Printf("registering player %d with name %q", i, s)
		pls.At(i).RegisterAs(s)
		i++
		return nil
	}
}

func (pls *Players) At(i int) *player.Player       { return pls.Players[i] }
func (pls *Players) All(prd player.Predicate) bool { return pls.Players.All(prd) }

type PlayedCards struct{ *set.Cards }

func (c PlayedCards) Pile() *set.Cards {
	if len(*c.Cards) == 5 {
		return (*set.Cards)(c.Cards)
	}
	return &set.Cards{}
}
