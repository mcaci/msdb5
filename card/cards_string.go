package card

func (cards Cards) String() string {
	var str string
	for _, cardID := range cards {
		c, _ := By(cardID)
		str += c.String() + " "
	}
	return str
}
