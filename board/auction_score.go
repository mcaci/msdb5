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
	currentScore, _ := strconv.Atoi(score)

	fun := func(x, m, r int) int {
		if currentScore*m <= x*m {
			return x * r
		}
		return currentScore
	}

	currentScore = fun(prevScore, 1, 1)
	currentScore = fun(minScore, 1, 1)
	currentScore = fun(maxScore, -1, 1)
	b.SetAuctionScore(uint8(currentScore))

	currentScore = fun(prevScore, 1, 0)
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
