package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type JoinStruct struct {
	request, origin string
}

func NewJoin(request, origin string) Action {
	return &JoinStruct{request, origin}
}
func NewJoinFinder(request, origin string) Finder {
	return &JoinStruct{request, origin}
}

func (js JoinStruct) Find(p *player.Player) bool { return isPlayerEmpty(p) }
func (js JoinStruct) Do(p *player.Player) error {
	name := strings.Split(js.request, "#")[1]
	p.Join(name, js.origin)
	return nil
}
func (js JoinStruct) NextPlayer(playerInTurn uint8) uint8 { return playersRoundRobin(playerInTurn) }
func (js JoinStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	if players.Count(predicate.NextPhasePlayerInfo) == 0 {
		return game.InsideAuction
	}
	return game.Joining
}
func (js JoinStruct) NextPhasePlayerInfo(p *player.Player) bool { return isPlayerEmpty(p) }

var isPlayerEmpty = func(p *player.Player) bool { return p.IsNameEmpty() }
