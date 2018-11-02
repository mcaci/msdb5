package card

// ID is the id of a card from 1 to 40
type ID uint8

// ID func
func (card *Card) ID() ID {
	return ID(card.number + (uint8)(card.seed)*10)
}

func toZeroBased(id ID) uint8 {
	return uint8(id) - 1
}

func fromZeroBased(index uint8) ID {
	return ID(index + 1)
}
