package msdb5

import "math"

var compareOn = func(a, b uint8) int { return int(a) - int(b) }

// Compare function
func (a Card) Compare(b Card) int {
	functionsToApply := []func(*Card) int{a.compareOnSeed, a.compareOnPoints, a.compareOnNumber}
	compareScore := 0
	for i := 0; i < len(functionsToApply) && compareScore == 0; i++ {
		compareScore = functionsToApply[i](&b)
	}
	return compareScore
}

func (a *Card) compareOnSeed(b *Card) int {
	return int(math.Abs(float64(a.seed) - float64(b.seed)))
}

func (a *Card) compareOnPoints(b *Card) int {
	return compareOn(a.points(), b.points())
}

func (a *Card) compareOnNumber(b *Card) int {
	return compareOn(a.number, b.number)
}
