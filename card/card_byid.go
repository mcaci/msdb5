package card

import (
	"errors"
)

// ByID func
func ByID(id uint8) (Card, error) {
	var card Card
	var err error
	if id < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if id > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		seedIndex := (id - 1) / 10
		number := intoNumber(id, seedIndex)
		seed := intoSeed(seedIndex)
		card = Card{number: number, seed: seed}
	}
	return card, err
}

func intoNumber(id, seedIndex uint8) uint8 {
	return id - (10 * seedIndex)
}

func intoSeed(seedIndex uint8) Seed {
	return Seed(seedIndex)
}
