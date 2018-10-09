package card

import (
	"errors"
	"strconv"
)

// Card type
type Card struct {
	number uint8
	seed   Seed
}

// ByName func
func ByName(number, seed string) (*Card, error) {
	var c Card
	var err error

	n, errN := strconv.Atoi(number)

	if errN != nil {
		err = errN
	} else if n > 10 || n < 1 {
		err = errors.New("number " + number + " doesn't exist")
	} else {
		c.number = uint8(n)
		if seed == Coin.String() {
			c.seed = Coin
		} else if seed == Cup.String() {
			c.seed = Cup
		} else if seed == Sword.String() {
			c.seed = Sword
		} else if seed == Cudgel.String() {
			c.seed = Cudgel
		} else {
			err = errors.New("seed " + seed + " doesn't exist")
		}
	}
	return &c, err
}

// ByID func
func ByID(id uint8) (*Card, error) {
	if id < 1 {
		return nil, errors.New("Index cannot be less than 1")
	} else if id > 40 {
		return nil, errors.New("Index cannot be more than 40")
	} else {
		seedIndex := (id - 1) / 10
		number := uint8(id - (10 * seedIndex))
		seed := Seed(seedIndex)
		return &Card{number: number, seed: seed}, nil
	}
}

// Number func
func (card *Card) Number() uint8 {
	return card.number
}

// Seed func
func (card *Card) Seed() Seed {
	return card.seed
}

// IsBriscola func
func (card *Card) IsBriscola(briscola Seed) bool {
	return card.seed == briscola
}

// Points func
func (card *Card) Points() uint8 {
	switch card.number {
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
