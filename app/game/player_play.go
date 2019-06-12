package game

import (
	"fmt"
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func processRequest(g *Game, rq *req, notify func(*player.Player, string)) error {
	err := play(g, *rq, notify)
	if err != nil {
		gamelog.ToConsole(g, g.sender(rq.From()), rq.Action(), err)
		notify(g.CurrentPlayer(), err.Error())
	}
	return err
}

func play(g *Game, rq req, notify func(*player.Player, string)) error {
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
