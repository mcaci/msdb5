package card

import (
	"errors"
	"strconv"
)

func ByName(number, seed string) (Card, error) {
	return HCard{number, seed}.Card()
}

// ByName func
func (hc HCard) Card() (Card, error) {
	var c Card
	var err error
	if c.number, err = hc.toNumber(); err == nil {
		c.seed, err = hc.toSeed()
	}
	return c, err
}

func (hc HCard) toNumber() (uint8, error) {
	n, err := strconv.Atoi(hc.number)

	if n > 10 || n < 1 {
		err = errors.New("hc.number '" + hc.number + "' doesn't exist")
	}
	return uint8(n), err
}

func (hc HCard) toSeed() (Seed, error) {
	var s Seed
	var err error

	switch hc.seed {
	case Coin.String():
		s = Coin
	case Cup.String():
		s = Cup
	case Sword.String():
		s = Sword
	case Cudgel.String():
		s = Cudgel
	default:
		err = errors.New("seed '" + hc.seed + "' doesn't exist")
	}
	return s, err
}
