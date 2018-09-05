package card

import "math"

// Compare function
func (card1 Card) Compare(card2 Card) int {
	compareFunctions := []func(*Card, *Card) int{compareOnSeed, compareOnPoints, compareOnNumber}
	info := cardsWithComparisonScoreInfo{card1: &card1, card2: &card2}
	for _, compareFunction := range compareFunctions {
		info.updateScore(compareFunction)
	}
	return info.score
}

type cardsWithComparisonScoreInfo struct {
	score        int
	card1, card2 *Card
}

func (c *cardsWithComparisonScoreInfo) updateScore(f func(*Card, *Card) int) int {
	if c.score == 0 {
		c.score = f(c.card1, c.card2)
	}
	return c.score
}

func compareOnSeed(card1, card2 *Card) int {
	return int(math.Abs(float64(card1.seed) - float64(card2.seed)))
}

func compareOnPoints(card1, card2 *Card) int { return compareOn(card1.points(), card2.points()) }

func compareOnNumber(card1, card2 *Card) int { return compareOn(card1.number, card2.number) }

func compareOn(card1, card2 uint8) int { return int(card1) - int(card2) }
