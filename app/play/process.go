package play

import (
	"container/list"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/nikiforosFreespirit/msdb5/app/msg"
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
	Lang() language.Tag
	Phase() phase.ID
	PlayedCards() *deck.Cards
	Players() team.Players
	SideDeck() *deck.Cards
}

type dataProvider interface {
	Value() string
	Action() string
	Card() (card.ID, error)
}

// Request func
func Request(g playInterface, rq dataProvider, setCompanion func(*player.Player), setBriscolaCard func(card.ID), sendMsg func(*player.Player, string)) error {
	p := g.CurrentPlayer()
	switch rq.Action() {
	case "Join":
		name := rq.Value()
		p.RegisterAs(name)
		return nil
	case "Auction":
		if p.Folded() {
			return nil
		}
		score, err := strconv.Atoi(rq.Value())
		if err != nil || !g.AuctionScore().CheckWith(auction.Score(score)) {
			p.Fold()
			return nil
		}
		g.AuctionScore().Update(auction.Score(score))
		if !g.IsSideUsed() {
			return nil
		}
		cardsToShow := SideCardsToDisplay(*g.AuctionScore())
		if cardsToShow == 0 {
			return nil
		}
		printer := message.NewPrinter(g.Lang())
		for _, pl := range g.Players() {
			printer.Fprintf(pl, "Side deck section: %s\n", msg.TranslateCards((*g.SideDeck())[:cardsToShow], printer))
		}
		return nil
	case "Exchange":
		if rq.Value() == "0" {
			return nil
		}
		c, err := rq.Card()
		if err != nil {
			return err
		}
		side := g.SideDeck()
		return Exchange(c, p.Hand(), side)
	case "Companion":
		c, err := rq.Card()
		if err != nil {
			return err
		}
		setBriscolaCard(c)
		_, pl := g.Players().Find(func(p *player.Player) bool { return p.Has(c) })
		setCompanion(pl)
		return nil
	case "Card":
		c, err := rq.Card()
		err = Play(p, c)
		if err != nil {
			return err
		}
		g.PlayedCards().Add(c)
		roundHasEnded := len(*g.PlayedCards()) == 5
		if roundHasEnded {
			playerIndex, _ := g.Players().Find(func(pl *player.Player) bool { return pl == p })
			winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
			var playersRoundRobin = func(playerIndex uint8) uint8 { return (playerIndex + 1) % 5 }
			next := playersRoundRobin(uint8(playerIndex) + winningCardIndex)
			g.Players()[next].Collect(g.PlayedCards())
			if team.Count(g.Players(), func(p *player.Player) bool { return p.IsHandEmpty() }) == 5 &&
				g.IsSideUsed() {
				side := g.SideDeck()
				g.Players()[next].Collect(side)
				side.Clear()
			}
		}
		return err
	default:
		return msg.ErrInvalidAction(rq.Action(), g.Lang())
	}
}
