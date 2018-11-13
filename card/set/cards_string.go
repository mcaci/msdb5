package set

func (cards Cards) String() string {
	var str string
	for _, cardID := range cards {
		str += cardID.String() + " "
	}
	return str
}
