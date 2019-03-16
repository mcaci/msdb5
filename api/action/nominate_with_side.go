package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type CompanionWithSideStruct struct {
	request, origin string
	playerInTurn    *player.Player
	players         playerset.Players
	set             func(card.ID, *player.Player)
}

func NewCompanionWithSide(request, origin string, playerInTurn *player.Player,
	players playerset.Players, set func(card.ID, *player.Player)) Action {
	return &CompanionWithSideStruct{request, origin, playerInTurn, players, set}
}

func (cs CompanionWithSideStruct) Phase() game.Phase { return game.ChosingCompanion }
func (cs CompanionWithSideStruct) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(cs.playerInTurn, cs.origin)
}
func (cs CompanionWithSideStruct) Do(p *player.Player) error {
	data := strings.Split(cs.request, "#")
	number := data[1]
	seed := data[2]
	c, err := card.Create(number, seed)
	if err != nil {
		return err
	}
	pl, err := cs.players.Find(func(p *player.Player) bool { return p.Has(c) })
	if err != nil {
		return err
	}
	cs.set(c, pl)
	return nil
}
func (cs CompanionWithSideStruct) NextPlayer(playerInTurn uint8) uint8 { return playerInTurn }
func (cs CompanionWithSideStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	return game.PlayingCards
}
func (cs CompanionWithSideStruct) NextPhasePlayerInfo(p *player.Player) bool { return true }
