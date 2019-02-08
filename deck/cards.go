package deck

import "github.com/nikiforosFreespirit/msdb5/card"

// Cards type
type Cards []card.ID

// Add func
func (cards *Cards) Add(ids ...card.ID) {
	*cards = append(*cards, ids...)
}

// Remove func
func (cards *Cards) Remove(index int) {
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}

// Has func
func (cards Cards) Has(id card.ID) bool {
	var found bool
	for _, cardID := range cards {
		if found = (cardID == id); found {
			break
		}
	}
	return found
}

// Supply func
func (cards *Cards) Supply() card.ID {
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
