package card

// ID is the id of a card from 1 to 40
type ID uint8

// Number func
func (id *ID) Number() uint8 {
	return id.idToNumber()
}

// Seed func
func (id *ID) Seed() Seed {
	return id.idToSeed()
}

