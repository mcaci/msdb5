package card

// Cards type
type Cards []uint8

// Add func
func (cards *Cards) Add(card Card) {
	*cards = append(*cards, card.ID())
}

// Has func
func (cards Cards) Has(card Card) bool {
	var cardFound bool
	for _, c := range cards {
		cardFound = (c == card.ID())
		if cardFound {
			break
		}
	}
	return cardFound
}

// FillWithIDs func
func FillWithIDs(ids ...uint8) Cards {
	var cards Cards
	for _, id := range ids {
		card, _ := ByID(id)
		cards.Add(card)
	}
	return cards
}
