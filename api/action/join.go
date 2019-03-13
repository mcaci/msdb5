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

func (js JoinStruct) Phase() game.Phase          { return game.Joining }
func (js JoinStruct) Find(p *player.Player) bool { return p.IsNameEmpty() }
func (js JoinStruct) Do(p *player.Player) error {
	data := strings.Split(js.request, "#")
	name := data[1]
	p.Join(name, js.origin)
	return nil
}
func (js JoinStruct) NextPlayer(playerInTurn uint8) uint8 { return nextPlayerInTurn(playerInTurn) }
func (js JoinStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) bool {
	return players.Count(predicate.NextPhasePlayerInfo) == 0
}
func (js JoinStruct) NextPhasePlayerInfo(p *player.Player) bool { return p.IsNameEmpty() }
