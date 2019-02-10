package card

import "strconv"

// ID is the id of a card from 1 to 40
type ID uint8

// Number func
func (id ID) Number() uint8 {
	return toZeroBased(id)%10 + 1
}

// Seed func
func (id ID) Seed() Seed {
	return Seed(toZeroBased(id) / 10)
}

func (id ID) String() string {
	return "(" + strconv.Itoa(int(id.Number())) + " of " + id.Seed().String() + ")"
}
