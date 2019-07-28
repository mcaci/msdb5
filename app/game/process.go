package game

import (
	"fmt"
	"os"

	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/request"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

// Process func
func (g *Game) Process(inputRequest, origin string) {
	printer := message.NewPrinter(g.Lang())
	rq := request.New(inputRequest, origin)
	report := func(err error) {
		fmt.Fprintf(os.Stdout, "New Action by %s: %s\nError raised: %+v\n", sender(g, rq).Name(), *rq, err)
		fmt.Fprintf(sender(g, rq), "Error: %+v\n", err)
	}

	// verify phase step
	err := request.VerifyPhase(g, rq)
	if err != nil {
		report(err)
		return
	}

	// verify player step
	err = request.VerifyPlayer(g, rq)
	if err != nil {
		report(err)
		return
	}

	// play step
	g.play(rq)

	if phase.InsideAuction == g.Phase() && len(*g.SideDeck()) != 0 {
		for _, pl := range g.Players() {
			fmt.Fprintf(pl, "Side deck section: %s\n", msg.TranslateCards((*g.SideDeck())[:auction.SideCards(*g.AuctionScore())], printer))
		}
	}

	// log action to file for ml
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}
	defer f.Close()
	// write to file for ml
	switch g.Phase() {
	case phase.ChoosingCompanion:
		fmt.Fprintf(f, "%s, %s, %d\n", g.CurrentPlayer().Name(), g.Companion().Name(), *(g.AuctionScore()))
	case phase.PlayingCards:
		lastPlayed := g.playedCards[len(g.playedCards)-1]
		fmt.Fprintf(f, "%s, %d\n", g.CurrentPlayer().Name(), lastPlayed)
	}
	// log action to console
	fmt.Fprintf(os.Stdout, "New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", sender(g, rq).Name(), *rq, sender(g, rq), g)

	// end round: next player
	plIndex := nextPlayer(g, rq)
	// next phase
	setCaller := func(p *player.Player) { g.caller = p }
	ph := nextPhase(g, rq, setCaller)
	fmt.Fprintf(g.LastPlayer(), msg.CreateInGameMsg(g, g.LastPlayer()))
	for _, pl := range g.Players() {
		printer.Fprintf(pl, "Game: %+v", msg.TranslateGameStatus(g, printer))
	}
	fmt.Fprintf(g.CurrentPlayer(), msg.CreateInGameMsg(g, g.CurrentPlayer()))
	// clean up
	g.cleanUp(plIndex)
	g.phase = ph
	track.Player(g.LastPlaying(), g.Players()[plIndex])

	if g.phase != phase.End {
		return
	}

	// process end phase
	remainingCards := len(*g.Players()[0].Hand())
	if remainingCards > 0 {
		collect(g)
	}
	// compute score
	pilers := make([]team.Piler, 0)
	for _, p := range g.Players() {
		pilers = append(pilers, p)
	}
	scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), pilers)
	for _, pl := range g.Players() {
		fmt.Fprintf(pl, "The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
	}
	// write to file
	fmt.Fprintf(f, "%s\n", g.CurrentPlayer().Name())
}

func sender(g *Game, rq requestInformer) *player.Player {
	index, _ := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(rq.From()) })
	return g.Players()[index]
}

type auctionData struct {
	pl    *player.Player
	score *auction.Score
}

func (a auctionData) Folded() bool                 { return player.Folded(a.pl) }
func (a auctionData) AuctionScore() *auction.Score { return a.score }
