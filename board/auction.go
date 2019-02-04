package board

import "strconv"

const minScore = 61
const maxScore = 120

// RaiseAuction func
func (b *Board) RaiseAuction(score, host string) {
	prevScore := int(b.AuctionScore())
	currentScore, _ := strconv.Atoi(score)
	currentScore = Compose(currentScore, NewAuction(prevScore, LT), NewAuction(minScore, LT), NewAuction(maxScore, GT))
	b.SetAuctionScore(uint8(currentScore))
	currentScore = Compose(currentScore, NewAuctionWithReturnScore(prevScore, 0, LT))
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
