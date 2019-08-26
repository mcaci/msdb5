package game

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/action"
	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/next"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

// Process func
func (g *Game) Process(inputRequest, origin string) []PlMsg {
	printer := message.NewPrinter(g.Lang())
	rq := NewReq(inputRequest)
	s := senderInfo{g.Players(), origin}
	r := report{}

	// verify phase step
	if r.err == nil {
		// err = msg.UnexpectedPhaseErr(phase.MustID(rq), g.Phase(), g.Lang())
		r.error(s, inputRequest, phase.Check(g, rq))
	}

	// verify player step
	if r.err == nil {
		// err = msg.UnexpectedPlayerErr(g.CurrentPlayer().Name(), g.Lang())
		es := expectedSenderInfo{s, g.CurrentPlayer()}
		r.error(es, inputRequest, team.CheckOrigin(es))
	}

	// play step
	if r.err == nil {
		r.error(s, inputRequest, action.Play(g, rq))
	}

	if r.err == nil {
		cardN := auction.SideCards(*g.AuctionScore())
		if phase.InsideAuction == g.Phase() && len(*g.SideDeck()) != 0 && cardN > 0 {
			for _, pl := range g.Players() {
				r.msg(pl, printer.Sprintf("Side deck section: (%s)\n", msg.TranslateCards((*g.SideDeck())[:cardN], printer)))
			}
		}

		// end round: next player
		plInfo := next.NewPlInfo(g.Phase(), g.Players(), g.PlayedCards(), g.Briscola(),
			len(*g.SideDeck()) > 0, len(*g.PlayedCards()) < 5, origin)
		nextPlayer := next.Player(plInfo)
		if g.Phase() == phase.PlayingCards && len(g.playedCards) == 5 {
			pile := nextPlayer.Pile()
			set.Move(g.PlayedCards(), pile)
			if team.Count(g.Players(), player.IsHandEmpty) == 5 && g.IsSideUsed() {
				set.Move(g.SideDeck(), pile)
			}
		}
		track.Player(g.LastPlaying(), nextPlayer)

		// end round: next phase
		phInfo := next.NewPhInfo(g.Phase(), g.Players(), g.Caller(), g.Companion(), g.Briscola(),
			len(*g.SideDeck()) > 0, len(*g.PlayedCards()) == 0, rq.Value())
		nextPhase := next.Phase(phInfo)
		if g.Phase() == phase.InsideAuction && nextPhase > g.Phase() {
			_, p := g.Players().Find(player.NotFolded)
			g.caller = p
		}
		g.setPhase(nextPhase)

		// log action to console
		senderPlayer := team.Sender(s)
		r.msg(os.Stdout, fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", senderPlayer.Name(), inputRequest, senderPlayer, g))
		for _, pl := range g.Players() {
			r.msg(pl, "-----")
		}
		r.msg(g.LastPlayer(), msg.CreateInGameMsg(g, g.LastPlayer()))
		for _, pl := range g.Players() {
			r.msg(pl, msg.TranslateGameStatus(g, printer))
		}
		r.msg(g.CurrentPlayer(), msg.CreateInGameMsg(g, g.CurrentPlayer()))

		if g.phase != phase.End {
			return r.reports
		}

		// process end phase
		remainingCards := len(*g.Players()[0].Hand())
		if remainingCards > 0 {
			highbriscolaCard := briscola.Serie(g.Briscola())
			for _, card := range highbriscolaCard {
				_, p := g.Players().Find(player.IsCardInHand(card))
				if p == nil { // no one has card
					continue
				}
				for _, pl := range g.Players() {
					set.Move(pl.Hand(), p.Pile())
				}
				if g.IsSideUsed() {
					set.Move(g.SideDeck(), p.Pile())
				}
				track.Player(g.LastPlaying(), p)
				printer := message.NewPrinter(g.Lang())
				team := printer.Sprintf("Callers")
				if p != g.Caller() && p != g.Companion() {
					team = printer.Sprintf("Others")
				}
				for _, pl := range g.Players() {
					r.msg(pl, printer.Sprintf("The end - %s team has all briscola cards", team))
				}
				break
			}
		}
		// compute score
		pilers := make([]team.Piler, 0)
		for _, p := range g.Players() {
			pilers = append(pilers, p)
		}
		scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), pilers, briscola.Points)
		for _, pl := range g.Players() {
			r.msg(pl, printer.Sprintf("The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2))
		}
	}
	r.msg(g.handleMLData()) // placeholder for ml data
	return r.reports
}

type report struct {
	reports []PlMsg
	err     error
}

func (r *report) msg(writer io.Writer, msg string) {
	r.reports = append(r.reports, PlMsg{writer, msg})
}

func (r *report) error(s team.SenderInformation, action string, err error) {
	r.msg(os.Stdout, fmt.Sprintf("New Action by %s: %s\nError raised: %+v\n", team.Sender(s).Name(), action, err))
	r.msg(team.Sender(s), fmt.Sprintf("Error: %+v\n", err))
	r.err = err
}

type expectedSenderInfo struct {
	senderInfo
	p *player.Player
}

func (s expectedSenderInfo) CurrentPlayer() *player.Player { return s.p }

type senderInfo struct {
	players team.Players
	origin  string
}

func (s senderInfo) From() string          { return s.origin }
func (s senderInfo) Players() team.Players { return s.players }
