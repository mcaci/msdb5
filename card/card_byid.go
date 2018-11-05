package card

import (
	"errors"
)

func ByID(id ID) (Card, error) {
	return id.ByID()
}

// ByID func
func (id ID) ByID() (Card, error) {
	var c Card
	var err error
	if c.number, err = id.extractNumber(); err == nil {
		c.seed, _ = id.extractSeed()
	}
	return c, err
}

func (id ID) extractNumber() (uint8, error) {
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

func (id ID) extractSeed() (Seed, error) {
	return Seed(toZeroBased(id) / 10), nil
}
