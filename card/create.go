package card

import (
	"errors"
	"strconv"
)

// Create func
func Create(number, seed string) (id ID, err error) {
	n, err := nameToNumber(number)
	if err != nil {
		return
	}
	s, err := nameToSeed(seed)
	if err != nil {
		return
	}
	index, err := mapToID(n, s)
	if err != nil {
		return
	}
	return ID(index), err
}

func mapToID(number uint8, seed Seed) (index uint8, err error) {
	index = number + (uint8)(seed)*10
	if index < 1 {
		err = errors.New("Index cannot be less than 1")
	}
	if index > 40 {
		err = errors.New("Index cannot be more than 40")
	}
	return
}

func nameToNumber(number string) (uint8, error) {
	n, err := strconv.Atoi(number)

	if n > 10 || n < 1 {
		err = errors.New("Number '" + number + "' is not valid for card")
	}
	return uint8(n), err
}

func nameToSeed(seed string) (Seed, error) {
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
