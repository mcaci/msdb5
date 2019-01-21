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

	LT := func(a, b int) bool { return a <= b }
	GT := func(a, b int) bool { return a >= b }
	fun := func(comp func(a, b int) bool, x, y, z int) int {
		if comp(x, y) {
			return z
		}
		return x
	}

	currentScore = fun(LT, currentScore, prevScore, prevScore)
	currentScore = fun(LT, currentScore, minScore, minScore)
	currentScore = fun(GT, currentScore, maxScore, maxScore)
	b.SetAuctionScore(uint8(currentScore))

	currentScore = fun(LT, currentScore, prevScore, 0)
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
