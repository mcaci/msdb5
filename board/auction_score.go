package board

import "strconv"

const minScore = 61
const maxScore = 120

// RaiseAuction func
func (b *Board) RaiseAuction(score string) {
	prevScore := int(b.AuctionScore())
	intScore, _ := strconv.Atoi(score)
	if intScore < minScore {
		intScore = minScore
	}
	if intScore <= prevScore {
		intScore = prevScore
	}
	if intScore > maxScore {
		intScore = maxScore
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
