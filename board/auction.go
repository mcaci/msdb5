package board

// AuctionCompare struct
type AuctionCompare struct {
	scoreToCompare, scoreToReturn int
	compareFunction               func(int, int) bool
}

// NewAuction func
func NewAuction(score int, compare func(int, int) bool) AuctionCompare {
	return NewAuctionWithReturnScore(score, score, compare)
}

// NewAuctionWithReturnScore func
func NewAuctionWithReturnScore(score, ret int, compare func(int, int) bool) AuctionCompare {
	return AuctionCompare{score, ret, compare}
}

// CompareAndAssignAuction func
func CompareAndAssignAuction(comp func(int, int) bool, x, y, z int) int {
	if comp(x, y) {
		return z
	}
	return x
}

// Compose func
func Compose(currentScore int, data ...AuctionCompare) int {
	for _, d := range data {
		currentScore = CompareAndAssignAuction(d.compareFunction, currentScore, d.scoreToCompare, d.scoreToReturn)
	}
	return currentScore
}

// LT var
var LT = func(a, b int) bool { return a <= b }

// GT var
var GT = func(a, b int) bool { return a >= b }
