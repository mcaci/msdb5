package card

import (
	"errors"
)

// ID is the id of a card from 1 to 40
type ID uint8

// ByID func
func ByID(id ID) (Card, error) {
	var card Card
	var err error
	if id < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if id > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		seedIndex := uint8(id-1) / 10
		number := intoNumber(id, seedIndex)
		seed := intoSeed(seedIndex)
		card = Card{number: number, seed: seed}
	}
	return card, err
}

func intoNumber(id ID, seedIndex uint8) uint8 {
	return uint8(id) - (10 * seedIndex)
}

func intoSeed(seedIndex uint8) Seed {
	return Seed(seedIndex)
}
