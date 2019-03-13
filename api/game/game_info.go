package game

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/display"
)

// Info func
func (g Game) Info() []display.Info {
	gameInfo := g.info.Info()
	plInTurn := display.Wrap("Turn of", g.PlayerInTurn().Name())
	gameInfo = append(gameInfo, plInTurn...)
	compCard := display.NewInfo("Companion", ":", g.companion.Card().String(), ";")
	gameInfo = append(gameInfo, compCard)
	return display.Wrap("Game", gameInfo...)
}

func (g Game) String() string {
	gameInfo := g.Info()
	players := display.NewInfo("Players", ":", g.players.String(), ";")
	gameInfo = append(gameInfo, players)
	if g.companion.Ref() != nil {
		gameInfo = append(gameInfo, g.companion.Ref().Name())
	}
	gameInfo = append(gameInfo, g.info.Info()...)
	phase := display.NewInfo("Phase", ":", strconv.Itoa(int(g.phase)), ";")
	gameInfo = append(gameInfo, phase)
	return display.All(display.Wrap("Game", gameInfo...)...)
}
