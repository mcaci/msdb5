package orchestrator

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

var playerSearchCriteria = func(g *Game, p *player.Player, origin string) bool {
	return p.IsRemoteHost(origin) && p == g.players[g.playerInTurn]
}

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
	playerLogged, _ := g.Players().Find(func(p *player.Player) bool { return p.IsRemoteHost(origin) })
	log.Printf("New Action by %s\n", playerLogged.Name())
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %v\n", err)
	log.Printf("Game info after action: %s\n", g.String())
	if g.phase == end {
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
	return g.Info(), g.players[g.playerInTurn].Info(), err
}

// Join func
func (g *Game) Join(name, origin string) (err error) {
	if g.phase != joining {
		err = errors.New("Phase is not joining")
	} else {
		nextPlayerJoining := func(p *player.Player) bool { return p.Name() == "" }
		p, err := g.Players().Find(nextPlayerJoining)
		if err == nil {
			p.Join(name, origin)
			if _, errNext := g.Players().Find(nextPlayerJoining); errNext != nil {
				g.phase = scoreAuction
				g.playerInTurn = 0
			}
		}
	}
	return err
}

// RaiseAuction func
func (g *Game) RaiseAuction(score, origin string) (err error) {
	if g.phase != scoreAuction {
		err = errors.New("Phase is not auction")
	} else {
		var p *player.Player
		p, err = g.Players().Find(func(p *player.Player) bool { return playerSearchCriteria(g, p, origin) })
		if err == nil {
			if !p.Folded() {
				prevScore := g.info.AuctionScore()
				currentScore, err := strconv.Atoi(score)
				if err != nil || uint8(currentScore) <= prevScore {
					p.Fold()
				} else {
					auction.Update(prevScore, uint8(currentScore), g.info.SetAuctionScore)
				}
			}
			nextPlayerIndex := (g.playerInTurn + 1) % 5
			for g.players[nextPlayerIndex].Folded() {
				nextPlayerIndex = (nextPlayerIndex + 1) % 5
			}
			g.playerInTurn = nextPlayerIndex
			foldCount := 0
			for _, pl := range g.players {
				if pl.Folded() {
					foldCount++
				}
			}
			if foldCount == 4 {
				g.phase = companionChoice
			}
		}
	}
	return
}

// Nominate func
func (g *Game) Nominate(number, seed, origin string) (err error) {
	if g.phase != companionChoice {
		err = errors.New("Phase is not auction")
	} else {
		_, err = g.Players().Find(func(p *player.Player) bool { return playerSearchCriteria(g, p, origin) })
		if err == nil {
			var c card.ID
			c, err = card.Create(number, seed)
			if err == nil {
				var p *player.Player
				p, err = g.Players().Find(func(p *player.Player) bool { return p.Has(c) })
				if err == nil {
					g.companion = *companion.New(c, p)
					g.phase = playBriscola
				}
			}
		}
	}
	return
}

// Play func
func (g *Game) Play(number, seed, origin string) (err error) {
	if g.phase != playBriscola {
		err = errors.New("Phase is not play")
	} else {
		var p *player.Player
		p, err = g.Players().Find(func(p *player.Player) bool { return playerSearchCriteria(g, p, origin) })
		if err == nil {
			var c card.ID
			c, err = p.Play(number, seed)
			if err == nil {
				roundHasEnded := g.info.PlayedCardIs(c)
				if roundHasEnded {
					playerIndex := (g.playerInTurn + briscola.IndexOfWinningCard(*g.info.PlayedCards(), g.companion.Card().Seed()) + 1) % 5
					g.info.PlayedCards().Move(g.Players()[playerIndex].Pile())
					g.playerInTurn = playerIndex
				} else {
					g.playerInTurn = (g.playerInTurn + 1) % 5
				}
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
		}
	}
	return
}
