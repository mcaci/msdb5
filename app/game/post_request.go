package game

import (
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
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

func postCompanionCard(cardProvider interface{ Card() card.ID },
	action interface{ SetBriscola(card.ID) }) {
	action.SetBriscola(cardProvider.Card())
}

func postCompanionPlayer(playerProvider interface{ Index() uint8 },
	action interface{ SetCompanion(uint8) }) {
	action.SetCompanion(playerProvider.Index())
}

func postExchangeCard(playerCardProvider interface {
	Card() card.ID
	Pl() *player.Player
},
	to interface{ SideDeck() *deck.Cards }) {

	cards := playerCardProvider.Pl().Hand()
	index := cards.Find(playerCardProvider.Card())
	toCards := to.SideDeck()
	awayCard := (*cards)[index]
	(*cards)[index] = (*toCards)[0]
	*toCards = append((*toCards)[1:], awayCard)
}

func postCardPlay(playerCardProvider interface {
	Card() card.ID
	Pl() *player.Player
},
	to interface{ PlayedCards() *deck.Cards }) {

	cards := playerCardProvider.Pl().Hand()
	index := cards.Find(playerCardProvider.Card())
	to.PlayedCards().Add((*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}
