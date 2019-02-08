package card

import (
	"errors"
	"strconv"
)

// ID is the id of a card from 1 to 40
type ID uint8

func byID(index uint8) (id ID, err error) {
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
