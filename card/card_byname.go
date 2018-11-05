package card

import (
	"errors"
	"strconv"
)

// ByName func
func ByName(number, seed string) (Card, error) {
	var c Card
	var err error
	if c.number, err = toNumber(number, seed); err == nil {
		c.seed, err = toSeed(number, seed)
	}
	return c, err
}

func toNumber(number, seed string) (uint8, error) {
	n, err := strconv.Atoi(number)

	if n > 10 || n < 1 {
		err = errors.New("number '" + number + "' doesn't exist")
	}
	return uint8(n), err
}

func toSeed(number, seed string) (Seed, error) {
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
