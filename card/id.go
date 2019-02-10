package card

import "strconv"

// ID is the id of a card from 1 to 40
type ID uint8

// Number func
func (id ID) Number() uint8 {
	return mapIDToNumber(id)
}

// Seed func
func (id ID) Seed() Seed {
	return mapIDToSeed(id)
}

func (id ID) String() string {
	return "(" + strconv.Itoa(int(id.Number())) + " of " + id.Seed().String() + ")"
}
