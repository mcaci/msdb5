package request

import (
	"container/list"

	"github.com/nikiforosFreespirit/msdb5/app/notify"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/app/track"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

type playerPredicate func(p *player.Player) bool

type expectedPlayerInterface interface {
	CurrentPlayer() *player.Player
	LastPlaying() *list.List
	Phase() phase.ID
	Players() team.Players
}

// requester
type requester interface {
	From() string
	Action() string
}

func FindCriteria(g expectedPlayerInterface, rq requester) playerPredicate {
	var expectedPlayerFinder playerPredicate
	switch rq.Action() {
	case "Join":
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsNameEmpty() }
	case "Origin":
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsSameHost(rq.From()) }
	default:
		expectedPlayerFinder = func(p *player.Player) bool { return p.IsExpectedPlayer(g.CurrentPlayer(), rq.From()) }
	}
	return expectedPlayerFinder
}

func VerifyPlayer(g expectedPlayerInterface, rq requester, sendMsg func(*player.Player, string)) error {
	criteria := FindCriteria(g, rq)
	_, actingPlayer, err := g.Players().Find(criteria)
	if err != nil {
		return notify.ErrPlayerNotFound(err, g.CurrentPlayer().Name())
	}
	if g.CurrentPlayer() == actingPlayer {
		return nil
	}
	track.Player(g.LastPlaying(), actingPlayer)
	return nil
}

func VerifyPhase(g expectedPlayerInterface, rq requester, sendMsg func(*player.Player, string)) error {
	currentPhase := g.Phase()
	inputPhase, err := phase.ToID(rq.Action())
	if err == nil && currentPhase == inputPhase {
		return nil
	}
	if err == nil && currentPhase != inputPhase {
		err = notify.ErrPhaseNotExpected(uint8(inputPhase), uint8(currentPhase))
	}
	return err
}
