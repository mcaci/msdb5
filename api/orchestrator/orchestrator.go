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
	find := isNameEmpty
	nextPlayerSupplier := func() uint8 { return 0 }
	nextPhasePredicate := func() bool { return g.players.Count(isNameEmpty) == 0 }
	return g.join(joining, name, origin, find, nextPlayerSupplier, nextPhasePredicate)
}
func (g *Game) join(phase phase, name, origin string, find func(*player.Player) bool, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	p, err := g.players.Find(find)
	if err != nil {
		return
	}
	p.Join(name, origin)
	if nextPhasePredicate() {
		g.nextPlayer(nextPlayerSupplier)
	}
	g.nextPhase(nextPhasePredicate)
	return err
}

// RaiseAuction func
func (g *Game) RaiseAuction(score, origin string) (err error) {
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	nextPlayerSupplier := func() uint8 { return 0 }
	nextPhasePredicate := func() bool { return g.players.Count(isNameEmpty) == 0 }
	return g.raiseAuction(scoreAuction, score, origin, find, nextPlayerSupplier, nextPhasePredicate)
}
func (g *Game) raiseAuction(phase phase, score, origin string, find func(*player.Player) bool, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	p, err := g.players.Find(find)
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
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	nextPlayerSupplier := func() uint8 { return g.playerInTurn }
	nextPhasePredicate := func() bool { return true }
	return g.nominate(companionChoice, number, seed, origin, find, nextPlayerSupplier, nextPhasePredicate)
}
func (g *Game) nominate(phase phase, number, seed, origin string, find func(*player.Player) bool, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	_, err = g.players.Find(find)
	if err != nil {
		return
	}
	c, err := card.Create(number, seed)
	if err != nil {
		return
	}
	p, err := g.players.Find(func(p *player.Player) bool { return p.Has(c) })
	if err != nil {
		return
	}
	g.setCompanion(c, p)
	g.nextPlayer(nextPlayerSupplier)
	g.nextPhase(nextPhasePredicate)
	return
}

// Play func
func (g *Game) Play(number, seed, origin string) (err error) {
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	var nextPlayerSupplier (func() uint8)
	nextPhasePredicate := func() bool { return verifyEndGame(g) }
	return g.play(playBriscola, number, seed, origin, find, nextPlayerSupplier, nextPhasePredicate)
}
func (g *Game) play(phase phase, number, seed, origin string, find func(*player.Player) bool, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	p, err := g.players.Find(find)
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
	g.nextPhase(nextPhasePredicate)
	return
}
