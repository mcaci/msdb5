package briscola5

type cmpInfo int8

const (
	LT_MIN_SCORE cmpInfo = iota - 2
	LE_ACTUAL
	GT_ACTUAL
	GE_MAX_SCORE
)

const (
	MIN_SCORE = 61
	MAX_SCORE = 120
)

// AuctionScore is the auction value
type AuctionScore uint8

// CmpAndSet compares two auction scores and returns the appropriate score
func CmpAndSet(actual, proposed AuctionScore) AuctionScore {
	switch Cmp(actual, proposed) {
	case LT_MIN_SCORE:
		return MIN_SCORE
	case LE_ACTUAL:
		return actual
	case GE_MAX_SCORE:
		return MAX_SCORE
	default:
		return proposed
	}
}

// Cmp compares two auction scores and returns the comparison information value
// -2 if both actual and proposed are less than 61
// -1 if proposed is less than actual but greater than 61
// 0 if proposed is greater than actual but less than 120
// 1 if proposed is greater than 120
func Cmp(actual, proposed AuctionScore) cmpInfo {
	switch {
	case proposed <= actual:
		return LE_ACTUAL
	case proposed < MIN_SCORE:
		return LT_MIN_SCORE
	case proposed >= MAX_SCORE:
		return GE_MAX_SCORE
	default:
		return GT_ACTUAL
	}
}
