package companion

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(s struct {
	ID      uint8
	Players team.Players
}) struct {
	Briscola  briscola.Card
	Companion uint8
} {
	count := make(map[card.Seed]uint8)
	for _, c := range *s.Players[s.ID].Hand() {
		count[c.Seed()]++
	}
	var max uint8
	var seed card.Seed
	for k, v := range count {
		if v < max {
			continue
		}
		seed = k
	}
	serie := briscola.Serie(Seeder(seed))
	for _, c := range serie {
		i, err := s.Players.Index(player.IsCardInHand(c))
		if err != nil {
			// card is in side deck
			continue
		}
		if i == s.ID {
			// player is self
			continue
		}
		return Round(briscola.Card{Item: c}, s.Players)
	}
	return Round(*briscola.MustID(1), s.Players)
}

type Seeder card.Seed

func (s Seeder) Seed() card.Seed { return card.Seed(s) }
