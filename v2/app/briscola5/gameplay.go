package briscola5

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola5/auction"
	"github.com/mcaci/msdb5/v2/app/briscola5/companion"
	"github.com/mcaci/msdb5/v2/app/briscola5/end"
	"github.com/mcaci/msdb5/v2/app/briscola5/exchange"
	"github.com/mcaci/msdb5/v2/app/briscola5/play"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Start(g *Game) {
	handSize := 7
	if !g.opts.WithSide {
		handSize++
	}
	// distribute cards to players
	briscola.Distribute(&struct {
		Players  briscola.Players
		Deck     *briscola.Deck
		HandSize int
	}{
		Players:  briscola.Players{Players: briscola5.ToGeneralPlayers(g.players)},
		Deck:     g.Deck(),
		HandSize: handSize,
	})
	// set side deck
	for range g.deck.Cards {
		g.side.Add(g.deck.Top())
	}

	// auction phase
	aucInf := auction.Run(g.players)
	g.auctionScore = aucInf.Score
	g.players.SetCaller(aucInf.Caller)

	// card exchange phase
	if g.opts.WithSide {
		exchange.Run(struct {
			Hand, Side *set.Cards
		}{
			Hand: g.players.Caller().Hand(),
			Side: &g.side.Cards,
		})
	}

	// companion choice phase
	cmpInf := companion.Run(
		struct {
			ID      uint8
			Players team.Players
		}{
			ID:      aucInf.Caller,
			Players: briscola5.ToGeneralPlayers(g.players),
		},
	)
	g.briscolaCard = cmpInf.Briscola
	g.players.SetCaller(cmpInf.Companion)

	// play phase
	plInfo := play.Run(struct {
		Players      briscola5.Players
		BriscolaCard briscola.Card
	}{
		Players:      g.players,
		BriscolaCard: cmpInf.Briscola,
	})

	// end phase
	end.Run(struct {
		PlayedCards  briscola.PlayedCards
		Players      team.Players
		BriscolaCard briscola.Card
		Side         briscola5.Side
	}{
		PlayedCards:  plInfo.OnBoard,
		Players:      briscola5.ToGeneralPlayers(g.players),
		BriscolaCard: cmpInf.Briscola,
		Side:         g.side,
	})
}
