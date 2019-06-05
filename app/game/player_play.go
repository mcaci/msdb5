package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func play(g *Game, request, origin string) error {
	p := g.CurrentPlayer()
	action := strings.Split(request, "#")[0]
	switch action {
	case "Join":
		name := strings.Split(request, "#")[1]
		p.RegisterAs(name)
		return nil
	case "Auction":
		score := strings.Split(request, "#")[1]
		currentScore, err := strconv.Atoi(score)
		if err == nil && g.auctionScore.CheckWith(auction.Score(currentScore)) && !p.Folded() {
			g.auctionScore.Update(auction.Score(currentScore))
		} else {
			p.Fold()
		}
		return nil
	case "Exchange":
		number := strings.Split(request, "#")[1]
		if number == "0" {
			return nil
		}
		seed := strings.Split(request, "#")[2]
		c, err := card.Create(number, seed)
		if err != nil {
			return err
		}
		return p.Exchange(c, &g.side)
	case "Companion":
		number := strings.Split(request, "#")[1]
		seed := strings.Split(request, "#")[2]
		c, err := card.Create(number, seed)
		if err != nil {
			return err
		}
		g.briscolaCard = c
		_, pl, err := g.players.Find(func(p *player.Player) bool { return p.Has(c) })
		if err == nil {
			g.companion = pl
		}
		return nil
	case "Card":
		number := strings.Split(request, "#")[1]
		seed := strings.Split(request, "#")[2]
		c, err := card.Create(number, seed)
		err = p.Play(c)
		if err != nil {
			return err
		}
		g.playedCards.Add(c)
		roundHasEnded := len(g.playedCards) == 5
		if roundHasEnded {
			playerIndex, _, _ := g.players.Find(func(pl *player.Player) bool { return pl == p })
			winningCardIndex := briscola.IndexOfWinningCard(g.playedCards, g.briscola())
			var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
			next := playersRoundRobin(uint8(playerIndex) + winningCardIndex)
			g.players[next].Collect(&g.playedCards)
			if team.Count(g.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5 &&
				len(g.side) > 0 {
				g.players[next].Collect(&g.side)
				g.side.Clear()
			}
		}
		return err
	default:
		return fmt.Errorf("Action %s not valid", action)
	}
}
