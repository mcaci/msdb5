package msdb5

func higherCardBetween(a, b *Card) *Card {
	c, haveSamePoints := compareOnPoints(a, b)
	if haveSamePoints {
		c = compareOnNumber(a, b)
	}
	return c
}

func compareOnPoints(a, b *Card) (*Card, bool) {
	var c *Card
	var haveSamePoints bool
	pointsForA := a.points()
	pointsForB := b.points()
	if pointsForA > pointsForB {
		c = a
	} else if pointsForB > pointsForA {
		c = b
	} else {
		haveSamePoints = true
	}
	return c, haveSamePoints
}

func compareOnNumber(a, b *Card) *Card {
	var c *Card
	if b.number >= a.number {
		c = b
	} else {
		c = a
	}
	return c
}
