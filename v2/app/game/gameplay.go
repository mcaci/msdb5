package game

import (
	"container/list"
	"fmt"
	"math/rand"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/game/auction"
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
	aucInf := auction.Run(&auction.Options{
		Players:     g.players,
		LastPlaying: g.lastPlaying,
	}, listen.WithTicker)
	g.auctionScore = aucInf.Score
	g.caller = aucInf.Caller

	// card exchange phase
	runExchange_v2(struct {
		opts *Options
		side set.Cards
		pl   *player.Player
	}{
		opts: g.opts,
		side: g.side,
		pl:   CurrentPlayer(g.lastPlaying),
	}, listen.WithTicker)

	// companion choice phase
	cmpInf := runCompanion_v2(struct {
		players team.Players
	}{
		players: g.players,
	}, func(id chan<- uint8) { id <- uint8(rand.Intn(40) + 1) })
	g.briscolaCard = *cmpInf.briscolaCard
	g.companion = cmpInf.companion

	// play phase
	plInfo := runPlay_v2(struct {
		players      team.Players
		briscolaCard card.Item
		lastPlaying  list.List
		caller       *player.Player
		companion    *player.Player
	}{
		players:      g.players,
		briscolaCard: *cmpInf.briscolaCard,
		lastPlaying:  g.lastPlaying,
		caller:       aucInf.Caller,
		companion:    cmpInf.companion})

	// end phase
	runEnd(struct {
		players      team.Players
		briscolaCard card.Item
		lastPlaying  list.List
		playedCards  set.Cards
		side         set.Cards
	}{
		players:      g.players,
		briscolaCard: *cmpInf.briscolaCard,
		lastPlaying:  g.lastPlaying,
		playedCards:  plInfo.onBoard,
	})
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
