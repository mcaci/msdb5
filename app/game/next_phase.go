package game

import (
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func nextPhase(g roundInformer, rq interface{ Value() string }) phase.ID {
	current := g.Phase()
	isNext := true
	switch current {
	case phase.Joining:
		isNext = team.Count(g.Players(), player.IsNameEmpty) == 0
	case phase.InsideAuction:
		isNext = team.Count(g.Players(), player.Folded) == 4
	case phase.ExchangingCards:
		isNext = rq.Value() == "0"
	case phase.PlayingCards:
		roundsBefore := uint8(len(*g.Players()[0].Hand()))
		isNext = predict(g, roundsBefore, 3) || checkAllWithEmptyHands(g)
	}
	if !isNext {
		return current
	}
	if current == phase.InsideAuction && !g.IsSideUsed() {
		return current + 2
	}
	return current + 1
}

func checkAllWithEmptyHands(g interface{ Players() team.Players }) bool {
	return team.Count(g.Players(), player.IsHandEmpty) == 5
}
