package card

import (
	"errors"
	"strconv"
)

// ID is the id of a card from 1 to 40
type ID uint8

// ByID func
func ByID(index uint8) (id ID, err error) {
	if index < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if index > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		id = ID(index)
	}
	return
}

// Number func
func (id *ID) Number() uint8 {
	return id.ToNumber()
}

// Seed func
func (id *ID) Seed() Seed {
	return id.ToSeed()
}

// ToNumber func
func (id ID) ToNumber() uint8 {
	return id.toZeroBased()%10 + 1
}

// ToSeed func
func (id ID) ToSeed() Seed {
	return Seed(id.toZeroBased() / 10)
}

func (id ID) toZeroBased() uint8 {
	return uint8(id) - 1
}

func (id ID) String() string {
	return "(" + strconv.Itoa(int(id.ToNumber())) + " of " + id.ToSeed().String() + ")"
}

// Points func
func (id ID) Points() uint8 {
	switch id.ToNumber() {
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
