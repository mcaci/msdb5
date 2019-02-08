package card

// Cards type
type Cards []ID

// Add func
func (cards *Cards) Add(ids ...ID) {
	*cards = append(*cards, ids...)
}

// Remove func
func (cards *Cards) Remove(index int) {
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}

// Has func
func (cards Cards) Has(id ID) bool {
	var found bool
	for _, cardID := range cards {
		if found = (cardID == id); found {
			break
		}
	}
	return found
}

// Supply func
func (cards *Cards) Supply() ID {
	card := (*cards)[0]
	(*cards) = (*cards)[1:]
	return card
}

// String func
func (cards Cards) String() string {
	var str string
	for _, cardID := range cards {
		str += cardID.String() + " "
	}
	return str
}
