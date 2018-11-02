package card

func (cards Cards) String() string {
	var str string
	for _, cardID := range cards {
		c, _ := ByID(cardID)
		str += c.String() + " "
	}
	return str
}
