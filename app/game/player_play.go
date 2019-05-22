package game

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func play(g *Game, p *player.Player, request, origin string) error {
	action := strings.Split(request, "#")[0]
	switch action {
	case "Join":
		name := strings.Split(request, "#")[1]
		p.Join(name, origin)
		return nil
	case "Auction":
		if p.Folded() {
			return nil
		}
		score := strings.Split(request, "#")[1]
		currentScore, err := strconv.Atoi(score)
		if err == nil && g.auctionScore.CheckWith(auction.Score(currentScore)) {
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
		if !p.Has(c) {
			return errors.New("card is not in players hand")
		}
		index, err := p.Hand().Find(c)
		if err != nil {
			return err
		}
		p.Hand().Add(g.side[0])
		g.side.Remove(0)
		g.side.Add(c)
		p.Hand().Remove(index)
		return nil
	case "Companion":
		number := strings.Split(request, "#")[1]
		seed := strings.Split(request, "#")[2]
		c, err := card.Create(number, seed)
		if err != nil {
			return err
		}
		_, pl, err := g.players.Find(func(p *player.Player) bool { return p.Has(c) })
		if err != nil {
			return err
		}
		g.SetCompanion(c, pl)
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
			playerInTurn, _, _ := g.players.Find(func(pl *player.Player) bool { return pl == p })
			winningCardIndex := briscola.IndexOfWinningCard(g.playedCards, g.BriscolaSeed())
			var playersRoundRobin = func(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }
			next := playersRoundRobin(uint8(playerInTurn) + winningCardIndex)
			g.players[next].Collect(g.PlayedCards())
			a := make([]player.EmptyHandChecker, 0)
			for _, p := range g.players {
				a = append(a, p)
			}
			if team.Count(g.players, func(p *player.Player) bool { return p.IsHandEmpty() }) == 5 &&
				len(g.side) > 0 {
				g.players[next].Collect(g.SideDeck())
				g.SideDeck().Clear()
			}
		}
		return err
	default:
		return fmt.Errorf("Action %s not valid", action)
	}
}
