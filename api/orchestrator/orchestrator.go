package orchestrator

import (
	"log"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
	"github.com/nikiforosFreespirit/msdb5/team"
)

// Orchestrator struct
type Orchestrator struct {
	game *game.Game
}

// NewAction func
func NewAction() api.Action {
	o := new(Orchestrator)
	o.game = game.NewGame()
	return o
}

var phaseActionMap map[string]func(*game.Game, string, string) action.Data

func init() {
	phaseActionMap = make(map[string]func(*game.Game, string, string) action.Data)
	phaseActionMap["Join"] = action.Join
	phaseActionMap["Auction"] = action.RaiseAuctionData
	phaseActionMap["Companion"] = action.NominateData
	phaseActionMap["Card"] = action.PlayData
}

// Action func
func (o *Orchestrator) Action(request, origin string) (all []display.Info, me []display.Info, err error) {
	inputAction := strings.Split(string(request), "#")[0]
	info := phaseActionMap[inputAction](o.game, request, origin)
	roundMayEnd := len(*o.game.Board().PlayedCards()) >= 4
	if roundMayEnd && o.game.CurrentPhase() == game.PlayBriscola {
		info = action.PlayEndRoundData(o.game, request, origin)
	}
	all, me, err = o.game.Info(), o.game.PlayerInTurn().Info(), playPhase(o.game, info)
	logEndRound(o.game, request, origin, err)
	if o.game.CurrentPhase() == game.End {
		all, me, err = endGame(o.game.Players(), o.game.Companion())
	}
	return
}

func endGame(players playerset.Players, companion player.ScoreCounter) ([]display.Info, []display.Info, error) {
	caller, _ := players.Find(func(p *player.Player) bool { return p.NotFolded() })
	team1, team2 := new(team.BriscolaTeam), new(team.BriscolaTeam)
	team1.Add(caller, companion)
	for _, pl := range players {
		if pl != caller && pl != companion {
			team2.Add(pl)
		}
	}
	return display.Wrap("Final Score", team1.Info("Callers"), team2.Info("Others")), nil, nil
}

func logEndRound(g *game.Game, request, origin string, err error) {
	playerLogged, _ := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(origin) })
	log.Printf("New Action by %s\n", playerLogged.Name().Display())
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %v\n", err)
	log.Printf("Game info after action: %s\n", g.String())
}
