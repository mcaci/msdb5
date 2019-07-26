package game

import (
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
)

func PostJoin(nameProvider interface{ Name() string },
	action interface{ RegisterAs(string) }) {
	action.RegisterAs(nameProvider.Name())
}

func PostAuctionFold(action interface{ Fold() }) {
	action.Fold()
}

func PostAuctionScore(scoreProvider interface{ Score() auction.Score },
	action interface{ SetAuction(auction.Score) }) {
	action.SetAuction(scoreProvider.Score())
}

func PostCompanionCard(cardProvider interface{ Card() card.ID },
	action interface{ SetBriscola(card.ID) }) {
	action.SetBriscola(cardProvider.Card())
}

func PostCompanionPlayer(playerProvider interface{ Index() uint8 },
	action interface{ SetCompanion(uint8) }) {
	action.SetCompanion(playerProvider.Index())
}

func PostExchange(cards, to *deck.Cards, index, toIndex int) {
	(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
}

func PostCardPlay(cards, to *deck.Cards, index int) {
	to.Add((*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}
