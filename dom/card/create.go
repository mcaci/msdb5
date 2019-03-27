package card

import (
	"errors"
	"strconv"
)

// Create func
func Create(number, seed string) (ID, error) {
	n, err := nameToNumber(number)
	if err != nil {
		return 0, err
	}
	s, err := nameToSeed(seed)
	return mapToID(n, s), err
}

func mapToID(number uint8, seed Seed) ID {
	return ID(number + (uint8)(seed)*10)
}

func nameToNumber(number string) (uint8, error) {
	n, err := strconv.Atoi(number)
	if n > 10 || n < 1 {
		err = errors.New("Number '" + number + "' is not valid for card")
	}
	return uint8(n), err
}

func nameToSeed(seed string) (s Seed, err error) {
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
	return
}
