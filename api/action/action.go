package action

import (
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func nextPlayer(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }

func isExpectedPlayer(p, other *player.Player, origin string) bool {
	return p.IsSame(other) && p.IsSameHost(origin)
}

func endGameCondition(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.All(searchCriteria)
}
