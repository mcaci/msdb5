package board

import "strconv"

const minScore = 61
const maxScore = 120

// RaiseAuction func
func (b *Board) RaiseAuction(score string) {
	prevScore := int(b.AuctionScore())
	intScore, err := strconv.Atoi(score)
	if err != nil || intScore <= prevScore {
		intScore = prevScore
	}
	if intScore < minScore {
		intScore = minScore
	}
	if intScore > maxScore {
		intScore = maxScore
	}
	b.SetAuctionScore(uint8(intScore))
}

// RaiseAuction2 func
func (b *Board) RaiseAuction2(score, host string) {
	prevScore := b.AuctionScore()
	intScore, _ := strconv.Atoi(score)
	currentScore := uint8(intScore)

	if currentScore <= prevScore {
		currentScore = prevScore
	}
	if currentScore < minScore {
		currentScore = minScore
	} else if currentScore > maxScore {
		currentScore = maxScore
	}
	b.SetAuctionScore(currentScore)

	if prevScore >= minScore && currentScore <= prevScore {
		currentScore = 0
	}
	p, _ := b.Players().Find(host)
	p.SetAuctionScore(currentScore)
}

// SetAuctionScore func
func (b *Board) SetAuctionScore(score uint8) {
	b.auctionScore = score
}

// AuctionScore func
func (b *Board) AuctionScore() uint8 {
	return b.auctionScore
}
