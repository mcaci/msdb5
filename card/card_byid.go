package card

import (
	"errors"
)

func ByID(id ID) (Card, error) {
	return id.Card()
}

// ByID func
func (id ID) Card() (Card, error) {
	var c Card
	var err error
	if c.number, err = id.toNumber(); err == nil {
		c.seed, _ = id.toSeed()
	}
	return c, err
}

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
