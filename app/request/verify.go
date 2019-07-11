package request

import (
	"container/list"
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/msg"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/app/track"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
	"golang.org/x/text/language"
)

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
		return msg.Error(fmt.Sprintf("Expecting player %s to play", g.CurrentPlayer().Name()), g.Lang())
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
	inputPhase, err := phase.ToID(rq.Action())
	if err == nil && currentPhase == inputPhase {
		return nil
	}
	if err == nil && currentPhase != inputPhase {
		err = msg.Error(fmt.Sprintf("Phase is not %d but %d", inputPhase, currentPhase), g.Lang())
	}
	return err
}
