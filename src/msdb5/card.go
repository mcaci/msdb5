package msdb5

import "strconv"

type Seed uint8

const (
	Coin Seed = iota
	Cup
	Sword
	Cudgel
)

type Card struct {
	number uint8
	seed   Seed
}

func (c *Card) points() uint8 {
	switch c.number {
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

func (card Card) String() string {
	return "(" + strconv.Itoa(int(card.number)) + " of " + card.seed.String() + ")"
}
