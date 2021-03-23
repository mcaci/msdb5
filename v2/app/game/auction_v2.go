package game

import (
	"context"
	"log"
	"math/rand"

	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func runAuction_v2(g *Game, listenFor func(context.Context, func())) struct {
	score  auction.Score
	caller *player.Player
} {
	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- 60 + rand.Intn(60) })
	go func() {
		<-done
		cancel()
		close(numbers)
	}()

	var caller *player.Player
	var scr auction.Score

	for score := range numbers {
		res := checkUpdateScore_v2(g, score)
		scr = res.score
		if res.fold != nil {
			res.fold()
		}
		if team.Count(g.players, notFolded) == 1 {
			// set caller
			caller = g.players.At(g.players.MustIndex(notFolded))
			// next phase
			g.phase++
			done <- struct{}{}
			close(done)
		}
		idx, err := CurrentPlayerIndex(CurrentPlayer(g.lastPlaying), g.players)
		if err != nil {
			log.Println(err)
		}
		// next player
		nextPlayer := roundRobin(idx, 1, numberOfPlayers)
		for player.Folded(g.players[nextPlayer]) {
			nextPlayer = roundRobin(nextPlayer, 1, numberOfPlayers)
		}
		track.Player(&g.lastPlaying, g.players[nextPlayer])
	}
	return struct {
		score  auction.Score
		caller *player.Player
	}{
		score:  scr,
		caller: caller,
	}
}

func rotate(start uint8, players team.Players) uint8 {
	nextPlayer := roundRobin(start, 1, numberOfPlayers)
	for player.Folded(players[nextPlayer]) {
		nextPlayer = roundRobin(nextPlayer, 1, numberOfPlayers)
	}
	return nextPlayer
}

func checkUpdateScore_v2(g *Game, score int) struct {
	score auction.Score
	fold  func()
} {
	toFold := player.Folded(CurrentPlayer(g.lastPlaying)) || !auction.ScoreCmp(g.auctionScore, auction.Score(score))
	if toFold {
		return struct {
			score auction.Score
			fold  func()
		}{
			score: g.auctionScore,
			fold:  func() { CurrentPlayer(g.lastPlaying).Fold() },
		}
	}
	newScore := auction.Max120(g.auctionScore, auction.Score(score))
	if newScore >= 120 {
		return struct {
			score auction.Score
			fold  func()
		}{
			score: newScore,
			fold: func() {
				for _, p := range g.players {
					if p == CurrentPlayer(g.lastPlaying) {
						continue
					}
					p.Fold()
				}
			},
		}
	}
	return struct {
		score auction.Score
		fold  func()
	}{
		score: g.auctionScore,
	}
}
