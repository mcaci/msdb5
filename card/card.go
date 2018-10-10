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
func ByName(number, seed string) (Card, error) {
	var c Card
	var err error
	c.number, err = toNumber(number)
	if err == nil {
		c.seed, err = toSeed(seed)
	}
	return c, err
}

func toNumber(number string) (uint8, error) {
	n, err := strconv.Atoi(number)

	if n > 10 || n < 1 {
		err = errors.New("number '" + number + "' doesn't exist")
	}
	return uint8(n), err
}

func toSeed(seed string) (Seed, error) {
	var s Seed
	var err error

	switch seed {
	case Coin.String():
		s = Coin
	case Cup.String():
		s = Cup
	case Sword.String():
		s = Sword
	case Cudgel.String():
		s = Cudgel
	default:
		err = errors.New("seed '" + seed + "' doesn't exist")
	}
	return s, err
}

// ByID func
func ByID(id uint8) (Card, error) {
	if id < 1 {
		return Card{}, errors.New("Index cannot be less than 1")
	} else if id > 40 {
		return Card{}, errors.New("Index cannot be more than 40")
	} else {
		seedIndex := (id - 1) / 10
		number := uint8(id - (10 * seedIndex))
		seed := Seed(seedIndex)
		return Card{number: number, seed: seed}, nil
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
