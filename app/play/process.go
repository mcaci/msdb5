package play

import (
	"container/list"
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
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
func Request(g playInterface, rq dataProvider, setCompanion func(*player.Player), setBriscolaCard func(card.ID)) error {
	p := g.CurrentPlayer()
	switch rq.Action() {
	case "Join":
		name := rq.Value()
		p.RegisterAs(name)
	case "Auction":
		if player.Folded(p) {
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
	case "Exchange":
		if rq.Value() == "0" {
			return nil
		}
		c, err := rq.Card()
		if err != nil {
			return err
		}
		return Exchange(c, p.Hand(), g.SideDeck())
	case "Companion":
		c, err := rq.Card()
		if err != nil {
			return err
		}
		return Companion(c, g.Players(), setCompanion, setBriscolaCard)
	case "Card":
		c, err := rq.Card()
		if err != nil {
			return err
		}
		err = Play(c, p.Hand())
		if err != nil {
			return err
		}
		g.PlayedCards().Add(c)
		roundHasEnded := len(*g.PlayedCards()) == 5
		if !roundHasEnded {
			break
		}
		playerIndex, _ := g.Players().Find(func(pl *player.Player) bool { return pl == p })
		winningCardIndex := briscola.IndexOfWinningCard(*g.PlayedCards(), g.Briscola())
		next := (playerIndex + int(winningCardIndex) + 1) % 5
		g.Players()[next].Collect(g.PlayedCards())
		if !(team.Count(g.Players(), player.IsHandEmpty) == 5 && g.IsSideUsed()) {
			break
		}
		side := g.SideDeck()
		g.Players()[next].Collect(side)
		side.Clear()
	default:
		return msg.Error(fmt.Sprintf("Action %s not valid", rq.Action()), g.Lang())
	}
	return nil
}
