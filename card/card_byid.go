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
		n = toZeroBased(id)%10 + 1
	}
	return n, err
}

func (id ID) toSeed() (Seed, error) {
	return Seed(toZeroBased(id) / 10), nil
}
