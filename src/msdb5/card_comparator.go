package msdb5

func (a Card) Compare(b Card) int {
	c := a.compareOnPoints(&b)
	if c == 0 {
		c = a.compareOnPoints(&b)
	}
	return c
}

func (a *Card) compareOnPoints(b *Card) int {
	pointsForA := a.points()
	pointsForB := b.points()
	return int(pointsForA - pointsForB)
}

func (a *Card) compareOnNumber(b *Card) int {
	return int(a.number - b.number)
}
