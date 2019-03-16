package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type CompanionStruct struct {
	request, origin string
	playerInTurn    *player.Player
	players         playerset.Players
	set             func(card.ID, *player.Player)
}

func NewCompanion(request, origin string, playerInTurn *player.Player,
	players playerset.Players, set func(card.ID, *player.Player)) Action {
	return &CompanionStruct{request, origin, playerInTurn, players, set}
}

func (cs CompanionStruct) Phase() game.Phase { return game.ChosingCompanion }
func (cs CompanionStruct) Find(p *player.Player) bool {
	return p.IsExpectedPlayer(cs.playerInTurn, cs.origin)
}
func (cs CompanionStruct) Do(p *player.Player) error {
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
func (cs CompanionStruct) NextPlayer(playerInTurn uint8) uint8 { return playerInTurn }
func (cs CompanionStruct) NextPhase(players playerset.Players, predicate PlayerPredicate) game.Phase {
	return game.PlayingCards
}
func (cs CompanionStruct) NextPhasePlayerInfo(p *player.Player) bool { return true }
