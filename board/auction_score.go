package board

import "strconv"

// RaiseAuction func
func (b *Board) RaiseAuction(score string) {
	prevScore := int(b.AuctionScore())
	intScore, _ := strconv.Atoi(score)
	if intScore < 61 {
		intScore = 61
	}
	if intScore <= prevScore {
		intScore = prevScore
	}
	if intScore > 120 {
		intScore = 120
	}
	b.SetAuctionScore(uint8(intScore))
}

// SetAuctionScore func
func (b *Board) SetAuctionScore(score uint8) {
	b.auctionScore = score
}

// AuctionScore func
func (b *Board) AuctionScore() uint8 {
	return b.auctionScore
}
