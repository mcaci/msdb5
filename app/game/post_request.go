package game

import (
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
)

func PostJoin(namer interface{ Name() string }, joiner interface{ CurrentPlayer() *player.Player }) {
	joiner.CurrentPlayer().RegisterAs(namer.Name())
}

func PostAuctionFold(auctioner interface{ CurrentPlayer() *player.Player }) {
	auctioner.CurrentPlayer().Fold()
}

func PostAuctionScore(scoreProvider interface{ Score() auction.Score },
	effector interface{ SetAuction(auction.Score) }) {
	effector.SetAuction(scoreProvider.Score())
}

func PostCompanionCard(cardProvider interface{ Card() card.ID },
	effector interface{ SetBriscola(card.ID) }) {
	effector.SetBriscola(cardProvider.Card())
}

func PostCompanionPlayer(playerProvider interface{ PlIdx() uint8 },
	effector interface{ SetCompanion(uint8) }) {
	effector.SetCompanion(playerProvider.PlIdx())
}

func PostExchange(cards, to *deck.Cards, index, toIndex int) {
	(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
}

func PostCardPlay(cards, to *deck.Cards, index int) {
	to.Add((*cards)[index])
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}
