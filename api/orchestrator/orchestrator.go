package orchestrator

import (
	"strings"

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
	logEndTurn(g, request, origin, err)
	infoForAllPlayers := g.Info()
	infoForSinglePlayer := g.players[g.playerInTurn].Info()
	if g.phase == end {
		infoForAllPlayers, infoForSinglePlayer, err = endGame(g)
	}
	return infoForAllPlayers, infoForSinglePlayer, err
}

// Join func
func (g *Game) Join(name, origin string) (err error) {
	err = phaseCheck(g, joining)
	if err == nil {
		nextPlayerJoining := func(p *player.Player) bool { return p.IsName("") }
		p, err := g.Players().Find(nextPlayerJoining)
		if err == nil {
			p.Join(name, origin)
			if _, errNext := g.Players().Find(nextPlayerJoining); errNext != nil {
				nextPhase(g, scoreAuction)
				nextPlayerTo(g, 0)
			}
		}
	}
	return err
}

// RaiseAuction func
func (g *Game) RaiseAuction(score, origin string) (err error) {
	err = phaseCheck(g, scoreAuction)
	if err == nil {
		var p *player.Player
		p, err = g.Players().Find(func(p *player.Player) bool { return playerSearchCriteria(g, p, origin) })
		if err == nil {
			updateAuction(g, p, score)
			nextAuctionPlayer(g)
			foldCount := countFoldedPlayers(g.players)
			if foldCount == 4 {
				nextPhase(g, companionChoice)
			}
		}
	}
	return
}

// Nominate func
func (g *Game) Nominate(number, seed, origin string) (err error) {
	err = phaseCheck(g, companionChoice)
	if err == nil {
		_, err = g.Players().Find(func(p *player.Player) bool { return playerSearchCriteria(g, p, origin) })
		if err == nil {
			var c card.ID
			c, err = card.Create(number, seed)
			if err == nil {
				var p *player.Player
				p, err = g.Players().Find(func(p *player.Player) bool { return p.Has(c) })
				if err == nil {
					setCompanion(g, c, p)
					nextPhase(g, playBriscola)
				}
			}
		}
	}
	return
}

// Play func
func (g *Game) Play(number, seed, origin string) (err error) {
	err = phaseCheck(g, playBriscola)
	if err == nil {
		var p *player.Player
		p, err = g.Players().Find(func(p *player.Player) bool { return playerSearchCriteria(g, p, origin) })
		if err == nil {
			var c card.ID
			c, err = p.Play(number, seed)
			if err == nil {
				roundHasEnded := verifyEndRound(g, c)
				if roundHasEnded {
					nextPlayerIndex := endRound(g)
					nextPlayerTo(g, nextPlayerIndex)
				} else {
					nextPlayer(g)
				}
				verifyEndGame(g)
			}
		}
	}
	return
}
