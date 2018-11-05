package card

// ID is the id of a card from 1 to 40
type ID uint8

// HCard is a card represented by string
type HCard struct {
	number, seed string
}

func toZeroBased(id ID) uint8 {
	return uint8(id) - 1
}

func fromZeroBased(index uint8) ID {
	return ID(index + 1)
}
