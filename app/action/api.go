package action

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type PhaseSupplier interface {
	ID() phase.ID
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
	NextPhase() phase.ID
}

type Cleaner interface {
	Clean()
}
