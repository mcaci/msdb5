package card

import "math"

// DoesOtherCardWin function
func DoesOtherCardWin(base, other *Card) bool {
	compareFunctions := []func(*Card, *Card) int{compareOnSeed, compareOnPoints, compareOnNumber}
	info := cardsWithComparisonScoreInfo{card1: base, card2: other}
	for _, compareFunction := range compareFunctions {
		info.updateScore(compareFunction)
	}
	return info.score < 0
}

type cardsWithComparisonScoreInfo struct {
	score        int
	card1, card2 *Card
}

func (c *cardsWithComparisonScoreInfo) updateScore(compare func(*Card, *Card) int) int {
	if c.score == 0 {
		c.score = compare(c.card1, c.card2)
	}
	return c.score
}

func compareOnSeed(card1, card2 *Card) int {
	return int(math.Abs(float64(card1.seed) - float64(card2.seed)))
}

func compareOnPoints(card1, card2 *Card) int { return compareOn(card1.Points(), card2.Points()) }

func compareOnNumber(card1, card2 *Card) int { return compareOn(card1.number, card2.number) }

func compareOn(card1, card2 uint8) int { return int(card1) - int(card2) }
