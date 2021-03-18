package game

import (
	"container/list"
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/listen"
	"github.com/mcaci/msdb5/v2/app/score"
	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/phase"
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
	runAuction(g, listen.WithTicker)

	// card exchange phase
	runExchange(g, listen.WithTicker)

	// companion choice phase
	runCompanion(g)

	// play phase
	runPlay_v2(struct {
		phase        phase.ID
		playedCards  set.Cards
		side         set.Cards
		players      team.Players
		briscolaCard card.Item
		lastPlaying  list.List
		caller       *player.Player
		companion    *player.Player
	}{phase: g.phase,
		playedCards:  g.playedCards,
		side:         g.side,
		players:      g.players,
		briscolaCard: g.briscolaCard,
		lastPlaying:  g.lastPlaying,
		caller:       g.caller,
		companion:    g.companion})

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
