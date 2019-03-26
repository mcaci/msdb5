package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/playerset"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
)

type JoinStruct struct {
	request, origin string
}

type JoinPhaseStruct struct {
	players playerset.Players
}

func NewJoin(request, origin string) Executer {
	return &JoinStruct{request, origin}
}
func NewPlayerSelector() NextPlayerSelector {
	return &JoinStruct{"", ""}
}
func NewJoinFinder(request, origin string) Finder {
	return &JoinStruct{request, origin}
}
func NewPhaseChanger(players playerset.Players) NextPhaseChanger {
	return &JoinPhaseStruct{players}
}

func (js JoinStruct) Find(p *player.Player) bool { return p.IsNameEmpty() }
func (js JoinStruct) Do(p *player.Player) error {
	name := strings.Split(js.request, "#")[1]
	p.Join(name, js.origin)
	return nil
}
func (js JoinStruct) NextPlayer(playerInTurn uint8) uint8 { return playersRoundRobin(playerInTurn) }
func (js JoinPhaseStruct) NextPhase() game.Phase {
	var isPlayerEmpty = func(p *player.Player) bool { return p.IsNameEmpty() }
	if js.players.Count(isPlayerEmpty) == 0 {
		return game.InsideAuction
	}
	return game.Joining
}
