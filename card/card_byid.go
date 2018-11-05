package card

import (
	"errors"
)

// ByID func
func ByID(id ID) (Card, error) {
	var c Card
	var err error
	if c.number, err = extractNumber(id); err == nil {
		c.seed, _ = extractSeed(id)
	}
	return c, err
}

func extractNumber(id ID) (uint8, error) {
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

func extractSeed(id ID) (Seed, error) {
	return Seed(toZeroBased(id) / 10), nil
}
