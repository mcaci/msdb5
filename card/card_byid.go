package card

import (
	"errors"
)

func (id ID) toNumber() (uint8, error) {
	var n uint8
	var err error
	if id < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if id > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		n = id.toZeroBased()%10 + 1
	}
	return n, err
}

func (id ID) toSeed() (Seed, error) {
	return Seed(id.toZeroBased() / 10), nil
}

// ToNumber func
func (id ID) ToNumber() uint8 {
	n, _ := id.toNumber()
	return n
}

// ToSeed func
func (id ID) ToSeed() Seed {
	s, _ := id.toSeed()
	return s
}

func (id ID) toZeroBased() uint8 {
	return uint8(id) - 1
}
