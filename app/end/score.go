package end

import (
	"github.com/nikiforosFreespirit/msdb5/app/notify"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
	"golang.org/x/text/language"
)

type endGameInformer interface {
	Caller() *player.Player
	Companion() *player.Player
	Lang() language.Tag
	Players() team.Players
}

// Score func
func Score(g endGameInformer, sendMsg func(*player.Player, string)) {
	scorers := make([]team.Scorer, 0)
	for _, p := range g.Players() {
		scorers = append(scorers, p)
	}
	scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), scorers...)
	for _, pl := range g.Players() {
		sendMsg(pl, notify.NotifyScore(scoreTeam1, scoreTeam2, g.Lang()))
	}
}
