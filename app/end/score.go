package end

import (
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type endGameInformer interface {
	Caller() *player.Player
	Companion() *player.Player
	Lang() language.Tag
	Players() team.Players
}

// Score func
func Score(g endGameInformer) {
	scorers := make([]team.Scorer, 0)
	for _, p := range g.Players() {
		scorers = append(scorers, p)
	}
	scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), scorers...)
	printer := message.NewPrinter(g.Lang())
	for _, pl := range g.Players() {
		printer.Fprintf(pl, "The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
	}
}
