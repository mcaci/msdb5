package card

import (
	"errors"
	"strconv"
)

// Create func
func Create(index uint8) (id ID, err error) {
	if index < 1 {
		err = errors.New("Index cannot be less than 1")
	} else if index > 40 {
		err = errors.New("Index cannot be more than 40")
	} else {
		id = ID(index)
	}
	return
}

// ByName func
func ByName(number, seed string) (id ID, err error) {
	var n uint8
	var s Seed
	if n, err = nameToNumber(number); err == nil {
		if s, err = nameToSeed(seed); err == nil {
			id, err = Create(mapToID(n, s))
		}
	}
	return
}

func nameToNumber(number string) (uint8, error) {
	n, err := strconv.Atoi(number)

	if n > 10 || n < 1 {
		err = errors.New("Number '" + number + "' doesn't exist")
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
