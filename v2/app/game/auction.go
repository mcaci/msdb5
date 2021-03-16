package game

import (
	"context"
	"math/rand"

	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

var notFolded player.Predicate = func(p *player.Player) bool { return !player.Folded(p) }

func runAuction(g *Game, listenFor func(context.Context, func())) {
	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- 60 + rand.Intn(60) })
	go func() {
		<-done
		cancel()
		close(numbers)
	}()

	for score := range numbers {
		checkUpdateScore(g, score)
		if team.Count(g.players, notFolded) == 1 {
			addAuctionInfo(g)
			done <- struct{}{}
			close(done)
		}
		nextPlayer(g)
	}
}

func nextPlayer(g *Game) {
	// next player
	nextPlayer := roundRobin(g.CurrentPlayerIndex(), 1, numberOfPlayers)
	for player.Folded(g.players[nextPlayer]) {
		nextPlayer = roundRobin(nextPlayer, 1, numberOfPlayers)
	}
	track.Player(&g.lastPlaying, g.players[nextPlayer])
}

func addAuctionInfo(g *Game) {
	// set caller
	g.caller = g.players.At(g.players.MustIndex(notFolded))
	// next phase
	g.phase++
}

func checkUpdateScore(g *Game, score int) {
	toFold := player.Folded(g.CurrentPlayer()) || !auction.CheckScores(g.auctionScore, auction.Score(score))
	if toFold {
		g.CurrentPlayer().Fold()
	}
	newScore := auction.Update(g.auctionScore, auction.Score(score))
	g.auctionScore = newScore
	if newScore >= 120 {
		for _, p := range g.players {
			if p == g.CurrentPlayer() {
				continue
			}
			p.Fold()
		}
	}
}
