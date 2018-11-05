package card

import (
	"errors"
)

// ByID func
func ByID(id ID) (Card, error) {
	var card Card
	var err error
	if id < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if id > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		number, _ := extractNumber(id)
		seed, _ := extractSeed(id)
		card = Card{number: number, seed: seed}
	}
	return card, err
}

func extractNumber(id ID) (uint8, error) {
	return toZeroBased(id)%10 + 1, nil
}

func extractSeed(id ID) (Seed, error) {
	return Seed(toZeroBased(id) / 10), nil
}
