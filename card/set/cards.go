package set

import "github.com/nikiforosFreespirit/msdb5/card"

// Cards type
type Cards []card.ID

// Add func
func (cards *Cards) Add(ids ...card.ID) {
	*cards = append(*cards, ids...)
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
