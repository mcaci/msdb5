package card

// Card type
type Card struct {
	number uint8
	seed   Seed
}

// ID func
func (card *Card) ID() ID {
	return ID(card.number + (uint8)(card.seed)*10)
}

// Number func
func (card *Card) Number() uint8 {
	return card.number
}

// Seed func
func (card *Card) Seed() Seed {
	return card.seed
}

// Points func
func (card *Card) Points() uint8 {
	switch card.number {
	case 1:
		return 11
	case 3:
		return 10
	case 8:
		return 2
	case 9:
		return 3
	case 10:
		return 4
	default:
		return 0
	}
}
