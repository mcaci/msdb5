package game

import (
	"fmt"
	"os"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/action"
	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/next"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/score"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

// Process func
func (g *Game) Process(inputRequest, origin string) []PlMsg {
	printer := message.NewPrinter(g.Lang())
	rq := NewReq(inputRequest)
	r := report{}

	// verify phase step
	if r.err == nil {
		// err = msg.UnexpectedPhaseErr(phase.MustID(rq), g.Phase(), g.Lang())
		s := senderInfo{g.Players(), origin}
		phInfo := phaseInfo{g.Phase(), rq.Action()}
		r.error(s, inputRequest, phase.Check(phInfo))
	}

	// verify player step
	if r.err == nil {
		// err = msg.UnexpectedPlayerErr(g.CurrentPlayer().Name(), g.Lang())
		s := senderInfo{g.Players(), origin}
		es := expectedSenderInfo{s, g.CurrentPlayer()}
		r.error(es, inputRequest, team.CheckOrigin(es))
	}

	// play step
	if r.err == nil {
		c, cerr := rq.Card()
		gInfo := gameRound{g, c, cerr, rq.Value()}
		s := senderInfo{g.Players(), origin}
		r.error(s, inputRequest, action.Play(gInfo))
	}

	if r.err != nil {
		return r.reports
	}

	if g.isToShow {
		for _, pl := range g.Players() {
			r.msg(pl, printer.Sprintf("Side deck section: (%s)\n", msg.TranslateCards(g.sideSubset, printer)))
		}
	}

	// end round: next player
	plInfo := next.NewPlInfo(g.Phase(), g.Players(), g.PlayedCards(), g.Briscola(),
		len(*g.SideDeck()) > 0, len(*g.PlayedCards()) < 5, origin)
	track.Player(g.LastPlaying(), next.Player(plInfo))
	if g.Phase() == phase.PlayingCards && len(g.playedCards) == 5 {
		set.Move(g.PlayedCards(), g.CurrentPlayer().Pile())
	}

	// end round: next phase
	phInfo := next.NewPhInfo(g.Phase(), g.Players(), g.Caller(), g.Companion(), g.Briscola(),
		len(*g.SideDeck()) > 0, len(*g.PlayedCards()) == 0, rq.Value())
	nextPhase := next.Phase(phInfo)
	if g.Phase() == phase.InsideAuction && nextPhase > g.Phase() {
		_, p := g.Players().Find(player.NotFolded)
		g.caller = p
	}
	g.setPhase(nextPhase)

	// send logs
	senderPlayer := team.Sender(senderInfo{g.Players(), origin})
	r.msg(os.Stdout, fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", senderPlayer.Name(), inputRequest, senderPlayer, g))
	for _, pl := range g.Players() {
		r.msg(pl, "-----")
	}
	r.msg(g.LastPlayer(), msg.CreateInGameMsg(g, g.LastPlayer()))
	gameStatusMsg := msg.TranslateGameStatus(g, printer)
	for _, pl := range g.Players() {
		r.msg(pl, gameStatusMsg)
	}
	r.msg(g.CurrentPlayer(), msg.CreateInGameMsg(g, g.CurrentPlayer()))

	if g.phase != phase.End {
		return r.reports
	}

	// process end game
	for _, card := range briscola.Serie(g.Briscola()) {
		_, p := g.Players().Find(player.IsCardInHand(card))
		if p == nil { // no one has card
			continue
		}
		team := printer.Sprintf("Callers")
		if p != g.Caller() && p != g.Companion() {
			team = printer.Sprintf("Others")
		}
		endMsg := printer.Sprintf("The end - %s team has all briscola cards", team)
		for _, pl := range g.Players() {
			r.msg(pl, endMsg)
		}
		track.Player(g.LastPlaying(), p)
		break
	}

	// last round winner collects all cards
	lastPlayerPile := g.CurrentPlayer().Pile()
	if len(*g.CurrentPlayer().Hand()) > 0 {
		for _, pl := range g.Players() {
			set.Move(pl.Hand(), lastPlayerPile)
		}
	}
	if g.IsSideUsed() {
		set.Move(g.SideDeck(), lastPlayerPile)
	}

	// compute score
	pilers := make([]score.Piler, len(g.Players()))
	for i, p := range g.Players() {
		pilers[i] = p
	}
	scoreTeam1, scoreTeam2 := score.Calc(g.Caller(), g.Companion(), pilers, briscola.Points)
	scoreMsg := printer.Sprintf("The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
	for _, pl := range g.Players() {
		r.msg(pl, scoreMsg)
	}
	r.msg(g.handleMLData()) // placeholder for ml data
	return r.reports
}
