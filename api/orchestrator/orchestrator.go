package orchestrator

import (
	"errors"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
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
	pInfo, err := g.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	return g.info.Info(), pInfo.Info(), err
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
		p, err = g.Players().Find(func(p *player.Player) bool { return p.Host() == origin && p == g.players[g.playerInTurn] })
		if err == nil {
			if !p.Folded() {
				prevScore := g.info.AuctionScore()
				currentScore, err := strconv.Atoi(score)
				if err != nil || uint8(currentScore) <= prevScore {
					p.Fold()
				} else {
					auction.Update(prevScore, prevScore, uint8(currentScore), g.info.SetAuctionScore)
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
		_, err = g.Players().Find(func(p *player.Player) bool { return p.Host() == origin && p == g.players[g.playerInTurn] })
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
func (g *Game) Play(number, seed, origin string) error {
	p, err := g.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	if err == nil {
		c, err := p.Play(number, seed)
		if err == nil {
			roundHasEnded := g.info.PlayedCardIs(c)
			if roundHasEnded {
				playerIndex := briscola.IndexOfWinningCard(*g.info.PlayedCards(), g.companion.Card().Seed())
				g.info.PlayedCards().Move(g.Players()[playerIndex].Pile())
			}
		}
	}
	return err
}
