package board

import "strconv"

// SetAuctionScore func
func (b *Board) RaiseAuction(score string) {
	scores, _ := strconv.Atoi(score)
	b.SetAuctionScore(uint8(scores))
}

// SetAuctionScore func
func (b *Board) SetAuctionScore(score uint8) {
	b.auctionScore = score
}

// AuctionScore func
func (b *Board) AuctionScore() uint8 {
	return b.auctionScore
}
