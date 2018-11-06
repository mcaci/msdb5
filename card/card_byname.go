package card

import (
	"errors"
	"strconv"
)

// ByName func
func ByName(number, seed string) (id ID, err error) {
	var n uint8
	var s Seed
	if n, err = toNumber(number); err == nil {
		if s, err = toSeed(seed); err == nil {
			id = ID(n + (uint8)(s)*10)
		}
	}
	return
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
