package board

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/board/auction"
)

const minScore = 61
const maxScore = 120

// RaiseAuction func
func (b *Board) RaiseAuction(score, host string) {
	prevScore := int(b.AuctionScore())
	currentScore, _ := strconv.Atoi(score)
	currentScore = auction.Compose(currentScore, auction.NewAuction(prevScore, auction.LT), auction.NewAuction(minScore, auction.LT), auction.NewAuction(maxScore, auction.GT))
	b.SetAuctionScore(uint8(currentScore))
	currentScore = auction.Compose(currentScore, auction.NewAuctionWithReturnScore(prevScore, 0, auction.LT))
	p, _ := b.Players().Find(host)
	p.SetAuctionScore(uint8(currentScore))
}

// SetAuctionScore func
func (b *Board) SetAuctionScore(score uint8) {
	b.auctionScore = score
}

// AuctionScore func
func (b *Board) AuctionScore() uint8 {
	return b.auctionScore
}
