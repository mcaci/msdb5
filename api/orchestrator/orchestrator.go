package orchestrator

import (
	"fmt"
	"log"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
	"github.com/nikiforosFreespirit/msdb5/team"
)

// Orchestrator struct
type Orchestrator struct {
	game *game.Game
}

// NewAction func
func NewAction(side bool) api.Action {
	o := new(Orchestrator)
	o.game = game.NewGame(side)
	return o
}

// Action func
func (o *Orchestrator) Action(request, origin string) (all, me string, err error) {
	data := strings.Split(request, "#")
	currentPlayer := o.game.PlayerInTurn()
	var actionExec action.Action
	switch data[0] {
	case "Join":
		actionExec = action.NewJoin(request, origin)
	case "Auction":
		actionExec = action.NewAuction(request, origin, currentPlayer,
			o.game.Players(), o.game.Board())
	case "Companion":
		actionExec = action.NewCompanion(request, origin, currentPlayer,
			o.game.Players(), o.game.SetCompanion)
	case "Card":
		actionExec = action.NewPlay(request, origin, o.game.PlayerInTurnIndex(),
			currentPlayer, o.game.Players(), o.game.Board(), o.game.BriscolaSeed())
	}
	err = playPhase(o.game, actionExec)
	all, me = fmt.Sprintf("Game: %+v", *o.game), fmt.Sprintf("%+v", currentPlayer)
	logEndRound(*o.game, request, origin, err)
	if o.game.CurrentPhase() == game.End {
		all, me, err = endGame(o.game.Players(), o.game.Companion())
	}
	return
}

func endGame(players playerset.Players, companion player.ScoreCounter) (string, string, error) {
	caller, _ := players.Find(func(p *player.Player) bool { return p.NotFolded() })
	team1, team2 := new(team.BriscolaTeam), new(team.BriscolaTeam)
	team1.Add(caller, companion)
	for _, pl := range players {
		if pl != caller && pl != companion {
			team2.Add(pl)
		}
	}
	return fmt.Sprintf("Callers: %+v; Others: %+v", team1.Score(), team2.Score()), "", nil
}

func logEndRound(g game.Game, request, origin string, err error) {
	playerLogged, err := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(origin) })
	if err == nil {
		log.Printf("New Action by %s\n", playerLogged.Name())
	}
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %+v\n", err)
	log.Printf("Game info after action: %+v\n", g.Log())
}
