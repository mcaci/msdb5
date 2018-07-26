package msdb5

func (a Card) Compare(b Card) int {
	c := a.compareOnPoints(&b)
	if c == 0 {
		c = a.compareOnNumber(&b)
	}
	return c
}

func (a *Card) compareOnPoints(b *Card) int {
	pointsForA := int(a.points())
	pointsForB := int(b.points())
	return pointsForA - pointsForB
}

func (a *Card) compareOnNumber(b *Card) int {
	numberForA := int(a.number)
	numberForB := int(b.number)
	return numberForA - numberForB
}
