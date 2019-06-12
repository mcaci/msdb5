package game

import (
	"fmt"
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func processRequest(g *Game, rq *req, notify func(*player.Player, string)) error {
	p := g.CurrentPlayer()
	switch rq.Action() {
	case "Join":
		name := rq.Value()
		p.RegisterAs(name)
		return nil
	case "Auction":
		score := rq.Value()
		currentScore, err := strconv.Atoi(score)
		if err == nil && g.auctionScore.CheckWith(auction.Score(currentScore)) && !p.Folded() {
			g.auctionScore.Update(auction.Score(currentScore))
			if g.IsSideUsed() {
				sideDeck := g.SideDeck()
				score := g.AuctionScore()
				if score >= 90 {
					for _, pl := range g.players {
						notify(pl, fmt.Sprintf("First card: %+v\n", sideDeck[0]))
					}
				}
				if score >= 100 {
					for _, pl := range g.players {
						notify(pl, fmt.Sprintf("Second card: %+v\n", sideDeck[1]))
					}
				}
				if score >= 110 {
					for _, pl := range g.players {
						notify(pl, fmt.Sprintf("Third card: %+v\n", sideDeck[2]))
					}
				}
				if score >= 120 {
					for _, pl := range g.players {
						notify(pl, fmt.Sprintf("Fourth and fifth cards: %+v, %+v\n", sideDeck[3], sideDeck[4]))
					}
				}
			}
		} else {
			p.Fold()
		}
		return nil
	case "Exchange":
		if rq.EndExchange() {
			return nil
		}
		c, err := rq.Card()
		if err != nil {
			return err
		}
		return p.Exchange(c, &g.side)
	case "Companion":
		c, err := rq.Card()
		if err != nil {
			return err
		}
		g.briscolaCard = c
		_, pl, err := g.players.Find(func(p *player.Player) bool { return p.Has(c) })
		if err != nil {
			return err
		}
		g.companion = pl
		return nil
	case "Card":
		c, err := rq.Card()
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
		return fmt.Errorf("Action %s not valid", rq.Action())
	}
}
