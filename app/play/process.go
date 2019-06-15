package play

import (
	"container/list"
	"fmt"
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

type playInterface interface {
	AuctionScore() *auction.Score
	Briscola() card.Seed
	CurrentPlayer() *player.Player
	IsSideUsed() bool
	LastPlaying() *list.List
	Phase() phase.ID
	PlayedCards() *deck.Cards
	Players() team.Players
	SideDeck() deck.Cards
}

type dataProvider interface {
	Value() string
	Action() string
	Card() (card.ID, error)
	EndExchange() bool
}

func Request(g playInterface, rq dataProvider, setCompanion func(*player.Player), setBriscolaCard func(card.ID), notify func(*player.Player, string)) error {
	p := g.CurrentPlayer()
	switch rq.Action() {
	case "Join":
		name := rq.Value()
		p.RegisterAs(name)
		return nil
	case "Auction":
		score := rq.Value()
		currentScore, err := strconv.Atoi(score)
		if err == nil && g.AuctionScore().CheckWith(auction.Score(currentScore)) && !p.Folded() {
			g.AuctionScore().Update(auction.Score(currentScore))
			if g.IsSideUsed() {
				sideDeck := g.SideDeck()
				score := *g.AuctionScore()
				if score >= 90 {
					for _, pl := range g.Players() {
						notify(pl, fmt.Sprintf("First card: %+v\n", sideDeck[0]))
					}
				}
				if score >= 100 {
					for _, pl := range g.Players() {
						notify(pl, fmt.Sprintf("Second card: %+v\n", sideDeck[1]))
					}
				}
				if score >= 110 {
					for _, pl := range g.Players() {
						notify(pl, fmt.Sprintf("Third card: %+v\n", sideDeck[2]))
					}
				}
				if score >= 120 {
					for _, pl := range g.Players() {
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
		side := g.SideDeck()
		return p.Exchange(c, &side)
	case "Companion":
		c, err := rq.Card()
		if err != nil {
			return err
		}
		setBriscolaCard(c)
		_, pl, err := g.Players().Find(func(p *player.Player) bool { return p.Has(c) })
		if err != nil {
			return err
		}
		setCompanion(pl)
		return nil
	case "Card":
		c, err := rq.Card()
		err = p.Play(c)
		if err != nil {
			return err
		}
		g.PlayedCards().Add(c)
		roundHasEnded := len(*g.PlayedCards()) == 5
		if roundHasEnded {
			playerIndex, _, _ := g.Players().Find(func(pl *player.Player) bool { return pl == p })
			winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
			var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
			next := playersRoundRobin(uint8(playerIndex) + winningCardIndex)
			g.Players()[next].Collect(g.PlayedCards())
			if team.Count(g.Players(), func(p *player.Player) bool { return p.IsHandEmpty() }) == 5 &&
				g.IsSideUsed() {
				side := g.SideDeck()
				g.Players()[next].Collect(&side)
				side.Clear()
			}
		}
		return err
	default:
		return fmt.Errorf("Action %s not valid", rq.Action())
	}
}
