package game

import (
	"container/list"
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/listen"
	"github.com/mcaci/msdb5/v2/app/score"
	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func WaitForPlayers(g *Game, listenFor func(chan<- string)) {
	names := make(chan string)
	go listenFor(names)
	for name := range names {
		p := player.New()
		p.RegisterAs(name)
		g.players = append(g.players, p)
	}
}

func Start(g *Game) {
	// distribute cards to players
	distributeCards(g)

	// setup first player
	track.Player(&g.lastPlaying, g.players[0])

	// auction phase
	aucInf := runAuction_v2(struct {
		players     team.Players
		lastPlaying list.List
	}{
		players:     g.players,
		lastPlaying: g.lastPlaying,
	}, listen.WithTicker)
	g.auctionScore = aucInf.score
	g.caller = aucInf.caller

	// card exchange phase
	runExchange(g, listen.WithTicker)

	// companion choice phase
	cmpInf := runCompanion_v2(struct {
		players team.Players
	}{
		players: g.players,
	})
	g.briscolaCard = *cmpInf.briscolaCard
	g.companion = cmpInf.companion

	// play phase
	runPlay_v2(struct {
		players      team.Players
		briscolaCard card.Item
		lastPlaying  list.List
		caller       *player.Player
		companion    *player.Player
	}{
		players:      g.players,
		briscolaCard: *cmpInf.briscolaCard,
		lastPlaying:  g.lastPlaying,
		caller:       aucInf.caller,
		companion:    cmpInf.companion})

	// end phase
	runEnd(g)
}

func Score(g *Game) string {
	t1, t2 := g.players.Part(team.IsInCallers(g))
	return fmt.Sprintf("[%s: %d], [%s: %d]",
		"Caller team", score.Sum(team.CommonPile(t1)),
		"Non Caller team", score.Sum(team.CommonPile(t2)))
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
