package game

import (
	"fmt"
	"math/rand"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/game/auction"
	"github.com/mcaci/msdb5/v2/app/game/companion"
	"github.com/mcaci/msdb5/v2/app/game/end"
	"github.com/mcaci/msdb5/v2/app/game/exchange"
	"github.com/mcaci/msdb5/v2/app/game/play"
	"github.com/mcaci/msdb5/v2/app/listen"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func WaitForPlayers(g *Game, listenFor func(chan<- string)) {
	names := make(chan string)
	go listenFor(names)
	for name := range names {
		p := briscola5.NewPlayer()
		p.RegisterAs(name)
		g.players = append(g.players, p)
	}
}

func Start(g *Game) {
	// distribute cards to players
	distributeCards(g)

	// auction phase
	aucInf := auction.Run(g.players, listen.WithTicker)
	g.auctionScore = aucInf.Score

	// card exchange phase
	if g.opts.WithSide {
		exchange.Run(struct {
			Hand, Side *set.Cards
		}{
			Hand: aucInf.Caller.Player.Hand(),
			Side: &g.side,
		}, listen.WithTicker)
	}

	// companion choice phase
	cmpInf := companion.Run(briscola5.ToGeneralPlayers(g.players), func(id chan<- uint8) { id <- uint8(rand.Intn(40) + 1) })
	g.briscolaCard = *cmpInf.Briscola
	g.cTeam = briscola5.NewCallersTeam(&aucInf.Caller.Player, cmpInf.Companion)

	// play phase
	plInfo := play.Run(struct {
		Players      team.Players
		BriscolaCard interface{ Seed() card.Seed }
		Callers      briscola5.Callerer
	}{
		Players:      briscola5.ToGeneralPlayers(g.players),
		BriscolaCard: *cmpInf.Briscola,
		Callers:      g.cTeam,
	})

	// end phase
	end.Run(struct {
		PlayedCards  set.Cards
		Players      team.Players
		BriscolaCard interface{ Seed() card.Seed }
		Side         set.Cards
	}{
		PlayedCards:  plInfo.OnBoard,
		Players:      briscola5.ToGeneralPlayers(g.players),
		BriscolaCard: *cmpInf.Briscola,
		Side:         g.side,
	})
}

func Score(g *Game) string {
	t1, t2 := briscola5.ToGeneralPlayers(g.players).Part(briscola5.IsInCallers(g.cTeam))
	return fmt.Sprintf("[%s: %d], [%s: %d]",
		"Caller team", briscola.Score(team.CommonPile(t1)),
		"Non Caller team", briscola.Score(team.CommonPile(t2)))
}

func distributeCards(g *Game) {
	d := set.Deck()
	for i := 0; i < set.DeckSize; i++ {
		if g.opts.WithSide && i >= set.DeckSize-5 {
			g.side.Add(d.Top())
			continue
		}
		g.players[i%5].Hand().Add(d.Top())
	}
}
