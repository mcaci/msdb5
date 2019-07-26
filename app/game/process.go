package game

import (
	"fmt"
	"os"

	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/deck"

	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/request"
	"github.com/mcaci/msdb5/app/track"
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
		printer.Fprintf(sender(g, rq), "Error: %+v\n", err)
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
	switch g.Phase() {
	case phase.Joining:
		data := phase.Join(rq)
		PostJoin(data, g)
	case phase.InsideAuction:
		data := phase.Auction(rq, auctionData{g.CurrentPlayer(), g.AuctionScore()})
		if data.ToFold() {
			PostAuctionFold(g)
			break
		}
		PostAuctionScore(data, g)
		if data.SideCards() == 0 {
			break
		}
		for _, pl := range g.Players() {
			printer.Fprintf(pl, "Side deck section: %s\n", msg.TranslateCards((*g.SideDeck())[:data.SideCards()], printer))
		}
	case phase.ExchangingCards:
		if rq.Value() == "0" {
			break
		}
		data := phase.Companion(rq, g.Players())
		plHand := g.players[data.PlIdx()].Hand()
		idx := plHand.Find(data.Card())
		func(cards, to *deck.Cards, index, toIndex int) {
			(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
		}(plHand, g.SideDeck(), idx, 0)
	case phase.ChoosingCompanion:
		data := phase.Companion(rq, g.Players())
		if err := data.CardNotFound(); err != nil {
			report(err)
			return
		}
		PostCompanionCard(data, g)
		PostCompanionPlayer(data, g)
	case phase.PlayingCards:
		data := phase.Companion(rq, g.Players())
		plHand := g.players[data.PlIdx()].Hand()
		idx := plHand.Find(data.Card())
		func(cards, to *deck.Cards, index, toIndex int) {
			to.Add((*cards)[index])
			*cards = append((*cards)[:index], (*cards)[index+1:]...)
		}(plHand, g.PlayedCards(), idx, 0)
	default:
		report(msg.Error(fmt.Sprintf("Action %s not valid", rq.Action()), g.Lang()))
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
	nextPlIdx := nextPlayer(g, rq)
	// next phase
	setCaller := func(p *player.Player) { g.caller = p }
	ph := nextPhase(g, rq, setCaller)
	// clean up
	cleanUp(g, rq)
	g.phase = ph
	track.Player(g.LastPlaying(), g.Players()[nextPlIdx])

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
		printer.Fprintf(pl, "The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
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
