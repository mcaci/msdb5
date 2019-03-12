package team

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// BriscolaTeam struct
type BriscolaTeam struct {
	players []player.ScoreCounter
}

// Add func
func (t *BriscolaTeam) Add(players ...player.ScoreCounter) {
	t.players = append(t.players, players...)
}

// Info func
func (t BriscolaTeam) Info(header string) display.Info {
	points := int(t.score(briscola.Points))
	return display.NewInfo(header, ":", strconv.Itoa(points), ";")
}

func (t BriscolaTeam) score(pointFunction func(card.ID) uint8) (score uint8) {
	for _, player := range t.players {
		score += player.Count(pointFunction)
	}
	return
}
