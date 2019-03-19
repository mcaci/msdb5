package action

import (
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type ExchangeCardsStruct struct {
	request, origin string
}

func NewExchangeCards(request, origin string) Action {
	return &ExchangeCardsStruct{request, origin}
}

func (ecs ExchangeCardsStruct) Phase() game.Phase                   { return game.ExchangingCards }
func (ecs ExchangeCardsStruct) Find(p *player.Player) bool          { return true }
func (ecs ExchangeCardsStruct) Do(p *player.Player) error           { return nil }
func (ecs ExchangeCardsStruct) NextPlayer(playerInTurn uint8) uint8 { return playerInTurn }
func (ecs ExchangeCardsStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	condition := true
	if condition {
		return game.ExchangingCards
	}
	return game.ChosingCompanion
}
func (ecs ExchangeCardsStruct) NextPhasePlayerInfo(p *player.Player) bool { return true }
