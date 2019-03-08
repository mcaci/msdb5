package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Action interface
func (g *Game) Action(request, origin string) ([]display.Info, []display.Info, error) {
	data := strings.Split(string(request), "#")
	var err error
	switch data[0] {
	case "Join":
		err = g.Join(data[1], origin)
	case "Auction":
		err = g.RaiseAuction(data[1], origin)
	case "Companion":
		err = g.Nominate(data[1], data[2], origin)
	case "Card":
		err = g.Play(data[1], data[2], origin)
	}
	logEndRound(g, request, origin, err)
	infoForAllPlayers := g.Info()
	infoForSinglePlayer := g.players[g.playerInTurn].Info()
	if g.phase == end {
		infoForAllPlayers, infoForSinglePlayer, err = endGame(g)
	}
	return infoForAllPlayers, infoForSinglePlayer, err
}

// Join func
func (g *Game) Join(name, origin string) (err error) {
	if err = g.phaseCheck(joining); err != nil {
		return
	}
	p, err := g.players.Find(isNameEmpty)
	if err != nil {
		return
	}
	p.Join(name, origin)
	if g.players.Count(isNameEmpty) == 0 {
		g.nextPlayer(func() uint8 { return 0 })
	}
	g.nextPhase(func() bool { return g.players.Count(isNameEmpty) == 0 })
	return err
}

// RaiseAuction func
func (g *Game) RaiseAuction(score, origin string) (err error) {
	if err = g.phaseCheck(scoreAuction); err != nil {
		return
	}
	p, err := g.players.Find(func(p *player.Player) bool { return isActive(g, p, origin) })
	if err != nil {
		return
	}
	auction.CheckAndUpdate(score, p.Folded, p.Fold, g.info.AuctionScore, g.info.SetAuctionScore)
	g.nextPlayer(func() uint8 {
		winnerIndex := (g.playerInTurn + 1) % 5
		for g.players[winnerIndex].Folded() {
			winnerIndex = (winnerIndex + 1) % 5
		}
		return winnerIndex
	})
	g.nextPhase(func() bool { return g.players.Count(folded) == 4 })
	return
}

// Nominate func
func (g *Game) Nominate(number, seed, origin string) (err error) {
	if err = g.phaseCheck(companionChoice); err != nil {
		return
	}
	if _, err = g.players.Find(func(p *player.Player) bool { return isActive(g, p, origin) }); err != nil {
		return
	}
	c, err := card.Create(number, seed)
	if err != nil {
		return
	}
	p, err := g.players.Find(func(p *player.Player) bool { return p.Has(c) })
	if err == nil {
		g.setCompanion(c, p)
	}
	g.nextPhase(func() bool { return err == nil })
	return
}

// Play func
func (g *Game) Play(number, seed, origin string) (err error) {
	if err = g.phaseCheck(playBriscola); err != nil {
		return
	}
	p, err := g.players.Find(func(p *player.Player) bool { return isActive(g, p, origin) })
	if err != nil {
		return
	}
	c, err := p.Play(number, seed)
	if err != nil {
		return
	}
	g.info.PlayedCards().Add(c)
	roundHasEnded := len(*g.info.PlayedCards()) >= 5
	if roundHasEnded {
		winnerIndex := winner(g)
		winnerPlayer := g.players[winnerIndex]
		winnerPlayer.Collect(g.info.PlayedCards())
		g.nextPlayer(func() uint8 { return winnerIndex })
	} else {
		g.nextPlayer(func() uint8 { return (g.playerInTurn + 1) % 5 })
	}
	g.nextPhase(func() bool { return verifyEndGame(g) })
	return
}
