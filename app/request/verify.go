package request

import (
	"container/list"
	"errors"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/language"
)

// ErrUnexpectedPlayer error
var ErrUnexpectedPlayer = errors.New("Unexpected player")

// ErrUnexpectedPhase error
var ErrUnexpectedPhase = errors.New("Unexpected phase")

type expectedPlayerFinder interface {
	CurrentPlayer() *player.Player
}

type requester interface {
	From() string
	Action() string
}

func FindCriteria(g expectedPlayerFinder, rq requester) player.Predicate {
	var expectedPlayerFinder player.Predicate
	switch rq.Action() {
	case "Join":
		expectedPlayerFinder = player.IsNameEmpty
	default:
		expectedPlayerFinder = func(p *player.Player) bool { return p == g.CurrentPlayer() && p.IsSameHost(rq.From()) }
	}
	return expectedPlayerFinder
}

type expectedPlayerInterface interface {
	CurrentPlayer() *player.Player
	LastPlaying() *list.List
	Lang() language.Tag
	Players() team.Players
}

func VerifyPlayer(g expectedPlayerInterface, rq requester) error {
	criteria := FindCriteria(g, rq)
	_, actingPlayer := g.Players().Find(criteria)
	if actingPlayer == nil {
		return ErrUnexpectedPlayer
	}
	if g.CurrentPlayer() == actingPlayer {
		return nil
	}
	track.Player(g.LastPlaying(), actingPlayer)
	return nil
}

type expectedPhaseInterface interface {
	Lang() language.Tag
	Phase() phase.ID
}

func VerifyPhase(g expectedPhaseInterface, rq requester) error {
	currentPhase := g.Phase()
	inputPhase, err := phase.ToID(rq)
	if err == nil && currentPhase == inputPhase {
		return nil
	}
	if err == nil && currentPhase != inputPhase {
		err = ErrUnexpectedPhase
	}
	return err
}
