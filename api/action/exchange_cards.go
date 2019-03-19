package action

import (
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type ExchangeCardsStruct struct {
	request, origin string
	playerInTurn    *player.Player
}

func NewExchangeCards(request, origin string, playerInTurn *player.Player) Action {
	return &ExchangeCardsStruct{request, origin, playerInTurn}
}
func (ecs ExchangeCardsStruct) Phase() game.Phase { return game.ExchangingCards }
func (ecs ExchangeCardsStruct) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(ecs.playerInTurn, ecs.origin)
}
func (ecs ExchangeCardsStruct) Do(p *player.Player) error           { return nil }
func (ecs ExchangeCardsStruct) NextPlayer(playerInTurn uint8) uint8 { return playerInTurn }
func (ecs ExchangeCardsStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	data := strings.Split(ecs.request, "#")
	number, err := strconv.Atoi(data[1])
	if number == 0 || err != nil {
		return game.ChosingCompanion
	}
	return game.ExchangingCards
}
func (ecs ExchangeCardsStruct) NextPhasePlayerInfo(p *player.Player) bool { return true }
