package card

import (
	"errors"
	"strconv"
)

// ByName func
func ByName(number, seed string) (ID, error) {
	var c Data
	var err error
	if c.number, err = toNumber(number); err == nil {
		c.seed, err = toSeed(seed)
	}
	return c.ID(), err
}

func toNumber(number string) (uint8, error) {
	n, err := strconv.Atoi(number)

	if n > 10 || n < 1 {
		err = errors.New("Number '" + number + "' doesn't exist")
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
		err = errors.New("Seed '" + seed + "' doesn't exist")
	}
	return s, err
}
