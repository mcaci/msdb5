package game

import (
	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

func nextPhase(g roundInformer, rq requestInformer, setCaller func(*player.Player)) phase.ID {
	current := g.Phase()
	nextPhase := current + 1
	switch {
	case current == phase.InsideAuction && !g.IsSideUsed():
		nextPhase = current + 2
	default:
		nextPhase = current + 1
	}
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
		isNext = predict(g, roundsBefore, 3) || check(g)
	}
	if isNext && current == phase.InsideAuction {
		_, p := g.Players().Find(func(p *player.Player) bool { return !player.Folded(p) })
		setCaller(p)
	}
	printer := message.NewPrinter(g.Lang())
	printer.Fprintf(g.LastPlayer(), msg.CreateInGameMsg(g, g.LastPlayer()))
	for _, pl := range g.Players() {
		printer.Fprintf(pl, "Game: %+v", g)
	}
	printer.Fprintf(g.CurrentPlayer(), msg.CreateInGameMsg(g, g.CurrentPlayer()))
	if !isNext {
		return current
	}
	return nextPhase
}
