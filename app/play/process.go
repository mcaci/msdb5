package play

import (
	"container/list"
	"fmt"

	"golang.org/x/text/language"

	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type playInterface interface {
	AuctionScore() *auction.Score
	Briscola() card.ID
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
	case "Exchange":
		if rq.Value() == "0" {
			return nil
		}
		return phase.CardAction(rq, p.Hand(), g.SideDeck(), func(cards, to *deck.Cards, index, toIndex int) {
			(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
		})
	case "Companion":
		c, err := rq.Card()
		if err != nil {
			return err
		}
		setBriscolaCard(c)
		_, pl := g.Players().Find(player.IsCardInHand(c))
		setCompanion(pl)
		return nil
	case "Card":
		return phase.CardAction(rq, p.Hand(), g.PlayedCards(), func(cards, to *deck.Cards, index, toIndex int) {
			to.Add((*cards)[index])
			*cards = append((*cards)[:index], (*cards)[index+1:]...)
		})
	default:
		return msg.Error(fmt.Sprintf("Action %s not valid", rq.Action()), g.Lang())
	}
	return nil
}
