package orchestrator

import (
	"errors"
	"log"
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/playerset"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

var playerSearchCriteria = func(g *Game, p *player.Player, origin string) bool {
	return p.IsRemoteHost(origin) && p == g.players[g.playerInTurn]
}

func logEndTurn(g *Game, request, origin string, err error) {
	playerLogged, _ := g.Players().Find(func(p *player.Player) bool { return p.IsRemoteHost(origin) })
	log.Printf("New Action by %s\n", playerLogged.Name())
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %v\n", err)
	log.Printf("Game info after action: %s\n", g.String())
}

func endGame(g *Game) ([]display.Info, []display.Info, error) {
	caller, _ := g.Players().Find(func(p *player.Player) bool { return !p.Folded() })
	score1 := caller.Count() + g.companion.Ref().Count()
	score2 := uint8(0)
	for _, pl := range g.Players() {
		if pl != caller && pl != g.companion.Ref() {
			score2 += pl.Count()
		}
	}
	score1info := display.NewInfo("Callers", ":", strconv.Itoa(int(score1)), ";")
	score2info := display.NewInfo("Others", ":", strconv.Itoa(int(score2)), ";")
	return display.Wrap("Final Score", score1info, score2info), nil, nil
}

func countFoldedPlayers(players playerset.Players) uint8 {
	foldCount := uint8(0)
	for _, pl := range players {
		if pl.Folded() {
			foldCount++
		}
	}
	return foldCount
}

func nextPhase(g *Game, next phase) {
	g.phase = next
}

func phaseCheck(g *Game, current phase) (err error) {
	if g.phase != current {
		err = errors.New("Phase is not " + strconv.Itoa(int(current)))
	}
	return
}

func setCompanion(g *Game, c card.ID, p *player.Player) {
	g.companion = *companion.New(c, p)
}

func updateAuction(g *Game, p *player.Player, score string) {
	if !p.Folded() {
		prevScore := g.info.AuctionScore()
		currentScore, err := strconv.Atoi(score)
		if err != nil || uint8(currentScore) <= prevScore {
			p.Fold()
		} else {
			auction.Update(prevScore, uint8(currentScore), g.info.SetAuctionScore)
		}
	}
}
func nextAuctionPlayer(g *Game) {
	nextPlayerIndex := (g.playerInTurn + 1) % 5
	for g.players[nextPlayerIndex].Folded() {
		nextPlayerIndex = (nextPlayerIndex + 1) % 5
	}
	g.playerInTurn = nextPlayerIndex
}

func nextPlayer(g *Game) {
	g.playerInTurn = (g.playerInTurn + 1) % 5
}

func nextPlayerToFirst(g *Game) {
	g.playerInTurn = 0
}

func startNewRound(g *Game) {
	playerIndex := (g.playerInTurn + briscola.IndexOfWinningCard(*g.info.PlayedCards(), g.companion.Card().Seed()) + 1) % 5
	g.info.PlayedCards().Move(g.Players()[playerIndex].Pile())
	g.playerInTurn = playerIndex
}

func verifyEndRound(g *Game, c card.ID) bool {
	return g.info.PlayedCardIs(c)
}

func verifyEndGame(g *Game) {
	gameHasEnded := true
	for _, pl := range g.players {
		if len(*pl.Hand()) > 0 {
			gameHasEnded = false
		}
	}
	if gameHasEnded {
		g.phase = end
	}
}
