package action

import (
	"container/list"
	"errors"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

// ErrUnexpectedPlayer error
var ErrUnexpectedPlayer = errors.New("Unexpected player")

// ErrUnexpectedPhase error
var ErrUnexpectedPhase = errors.New("Unexpected phase")

type requester interface {
	From() string
	Action() string
}

func findCriteria(g interface{ CurrentPlayer() *player.Player }, rq requester) player.Predicate {
	var crit player.Predicate
	switch rq.Action() {
	case "Join":
		crit = player.IsNameEmpty
	default:
		crit = func(p *player.Player) bool { return p == g.CurrentPlayer() && p.IsSameHost(rq.From()) }
	}
	return crit
}

type expectedPlayerInterface interface {
	CurrentPlayer() *player.Player
	LastPlaying() *list.List
	Players() team.Players
}

func VerifyPlayer(g expectedPlayerInterface, rq requester) error {
	criteria := findCriteria(g, rq)
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

func VerifyPhase(g interface{ Phase() phase.ID }, rq requester) error {
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
