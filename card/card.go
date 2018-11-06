package card

import "errors"

// Data type
type Data struct {
	number uint8
	seed   Seed
}

// ID is the id of a card from 1 to 40
type ID uint8

// Card func
func Card(index uint8) (id ID, err error) {
	if index < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if index > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		id = ID(index)
	}
	return
}

// ID func
func (card *Data) ID() ID {
	return ID(card.number + (uint8)(card.seed)*10)
}

// Number func
func (card *ID) Number() uint8 {
	return card.ToNumber()
}

// Seed func
func (card *ID) Seed() Seed {
	return card.ToSeed()
}

// Points func
func (card *ID) Points() uint8 {
	switch card.ToNumber() {
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
