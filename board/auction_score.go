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
	if prevScore >= minScore && intScore <= prevScore {
		p, _ := b.Players().Find(host)
		p.SetAuctionScore(0)
	} else {
		p, _ := b.Players().Find(host)
		p.SetAuctionScore(uint8(intScore))
	}
}

// SetAuctionScore func
func (b *Board) SetAuctionScore(score uint8) {
	b.auctionScore = score
}

// AuctionScore func
func (b *Board) AuctionScore() uint8 {
	return b.auctionScore
}
