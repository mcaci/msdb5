package end

import (
	"container/list"
	"io"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/card"

	"github.com/nikiforosFreespirit/msdb5/app/notify"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/app/track"
	_ "github.com/nikiforosFreespirit/msdb5/catalog"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

type playersInformer interface {
	Caller() *player.Player
	Companion() *player.Player
	Players() team.Players
	Briscola() card.Seed
	LastPlaying() *list.List
	Lang() language.Tag
	CardsOnTheBoard() int
	IsSideUsed() bool
	SideDeck() *deck.Cards
}

func Check(g playersInformer, sendMsg func(*player.Player, string)) bool {
	roundsLeft := len(*g.Players()[0].Hand())
	if g.CardsOnTheBoard() >= 5 && roundsLeft <= 3 {
		highbriscolaCard := deck.Highest(g.Briscola())
		var callers, others bool
		var roundsChecked int
		for _, card := range highbriscolaCard {
			if roundsChecked == roundsLeft {
				break
			}
			_, p := g.Players().Find(func(p *player.Player) bool {
				return p.Has(card)
			})
			if p == nil { // no one has card
				continue
			}
			if p == g.Caller() || p == g.Companion() {
				callers = true
			} else {
				others = true
			}
			if callers == others {
				break
			}
			roundsChecked++
		}
		if callers != others {
			p := g.Caller()
			printer := message.NewPrinter(g.Lang())
			team := printer.Sprintf("Callers")
			if others {
				_, p = g.Players().Find(func(p *player.Player) bool {
					return p == g.Caller() || p == g.Companion()
				})
				team = printer.Sprintf("Others")
			}
			collect(g, p, team, sendMsg)

			return true
		}
	}
	return team.Count(g.Players(), func(p *player.Player) bool { return p.IsHandEmpty() }) == 5
}

func collect(g playersInformer, p *player.Player, team string, sendMsg func(*player.Player, string)) {
	for _, pl := range g.Players() {
		p.Collect(pl.Hand())
		sendMsg(pl, notify.NotifyAnticipatedEnding(team, g.Lang()))
	}
	if g.IsSideUsed() {
		p.Collect(g.SideDeck())
	}
	track.Player(g.LastPlaying(), p)
}

type endGameInformer interface {
	AuctionScore() *auction.Score
	Caller() *player.Player
	Companion() *player.Player
	CurrentPlayer() *player.Player
	Lang() language.Tag
	LastCardPlayed() card.ID
	Phase() phase.ID

	Players() team.Players
	// not registerd yet
	IsSideUsed() bool
	SideDeck() *deck.Cards
}

// Process func
func Process(g endGameInformer, file io.Writer, sendMsg func(*player.Player, string)) error {
	scorers := make([]team.Scorer, 0)
	for _, p := range g.Players() {
		scorers = append(scorers, p)
	}
	scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), scorers...)
	for _, pl := range g.Players() {
		sendMsg(pl, notify.NotifyScore(scoreTeam1, scoreTeam2, g.Lang()))
	}
	notify.ToFile(g, file)
	return nil
}
