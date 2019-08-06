package game

import (
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
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

func postExchange(cards, to *deck.Cards, index int) {
	awayCard := (*cards)[index]
	(*cards)[index] = (*to)[0]
	*to = append((*to)[1:], awayCard)
}

func postCardPlay(cards, to *deck.Cards, index int) {
	to.Add((*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}
