package game

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
)

func postJoin(nameProvider interface{ Name() string },
	action interface{ RegisterAs(string) }) {
	action.RegisterAs(nameProvider.Name())
}

func postAuctionFold(action interface{ Fold() }) {
	action.Fold()
}

func postAuctionScore(scoreProvider interface{ Score() auction.Score },
	action interface{ SetAuction(auction.Score) }) {
	action.SetAuction(scoreProvider.Score())
}

type playerCardProvider interface {
	Card() *card.Item
	Pl() *player.Player
}

func postCompanion(plCProv playerCardProvider, action interface {
	SetBriscola(*card.Item)
	SetCompanion(*player.Player)
}) {
	action.SetBriscola(plCProv.Card())
	action.SetCompanion(plCProv.Pl())
}

func postExchange(plCProv playerCardProvider, to interface{ SideDeck() *set.Cards }) {
	cards := plCProv.Pl().Hand()
	index := cards.Find(*plCProv.Card())
	toCards := to.SideDeck()
	awayCard := (*cards)[index]
	(*cards)[index] = (*toCards)[0]
	*toCards = append((*toCards)[1:], awayCard)
}

func postPlay(plCProv playerCardProvider, to interface{ PlayedCards() *set.Cards }) {
	cards := plCProv.Pl().Hand()
	index := cards.Find(*plCProv.Card())
	to.PlayedCards().Add((*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}
