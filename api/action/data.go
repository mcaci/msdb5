package action

import (
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

// Data struct
type Data struct {
	phase              game.Phase
	find               func(*player.Player) bool
	do                 func(*player.Player) error
	nextPlayerOperator func(uint8) uint8
	nextPhasePredicate func(playerset.Players, func(*player.Player) bool) bool
	playerPredicate    func(*player.Player) bool
}

func (d Data) Phase() game.Phase                          { return d.phase }
func (d Data) Find() func(*player.Player) bool            { return d.find }
func (d Data) Do() func(*player.Player) error             { return d.do }
func (d Data) NextPlayerOperator() func(uint8) uint8      { return d.nextPlayerOperator }
func (d Data) PlayerPredicate() func(*player.Player) bool { return d.playerPredicate }
func (d Data) NextPhasePredicate() func(playerset.Players, func(*player.Player) bool) bool {
	return d.nextPhasePredicate
}
