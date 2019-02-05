package auction

// AuctionComparisonData struct
type AuctionComparisonData struct {
	scoreToCompare, scoreToReturn int
	compareFunction               func(int, int) bool
}

// NewAuction func
func NewAuction(score int, compare func(int, int) bool) *AuctionComparisonData {
	return NewAuctionWithReturnScore(score, score, compare)
}

// NewAuctionWithReturnScore func
func NewAuctionWithReturnScore(score, ret int, compare func(int, int) bool) *AuctionComparisonData {
	return &AuctionComparisonData{score, ret, compare}
}

// Compose func
func Compose(currentScore int, data ...*AuctionComparisonData) int {
	for _, d := range data {
		currentScore = d.compareAndAssign(currentScore)
	}
	return currentScore
}

// LT var
var LT = func(a, b int) bool { return a <= b }

// GT var
var GT = func(a, b int) bool { return a >= b }

func (data *AuctionComparisonData) compareAndAssign(currentScore int) int {
	if data.compareFunction(currentScore, data.scoreToCompare) {
		return data.scoreToReturn
	}
	return currentScore
}
