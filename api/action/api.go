package action

import (
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
)

type PhaseSupplier interface {
	Phase() game.Phase
}

type Finder interface {
	Find(*player.Player) bool
}

type Executer interface {
	Do(*player.Player) error
}

type NextPlayerSelector interface {
	NextPlayer(uint8) uint8
}

type NextPhaseChanger interface {
	NextPhase() game.Phase
}
