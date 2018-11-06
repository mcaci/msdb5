package card

import (
	"errors"
	"strconv"
)

func (sCard StrData) toNumber() (uint8, error) {
	n, err := strconv.Atoi(sCard.number)

	if n > 10 || n < 1 {
		err = errors.New("Number '" + sCard.number + "' doesn't exist")
	}
	return uint8(n), err
}

func (sCard StrData) toSeed() (Seed, error) {
	var s Seed
	var err error

	switch sCard.seed {
	case Coin.String():
		s = Coin
	case Cup.String():
		s = Cup
	case Sword.String():
		s = Sword
	case Cudgel.String():
		s = Cudgel
	default:
		err = errors.New("Seed '" + sCard.seed + "' doesn't exist")
	}
	return s, err
}
