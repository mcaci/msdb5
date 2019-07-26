package card

import "fmt"

// ID is the id of a card from 1 to 40
type ID uint8

// Number func
func (id ID) Number() uint8 {
	return (uint8(id)-1)%10 + 1
}

// Seed func
func (id ID) Seed() Seed {
	return Seed((uint8(id) - 1) / 10)
}

func (id ID) String() string {
	if id == 0 {
		return "(Undefined card)"
	}
	seeds := []string{"Coin", "Cup", "Sword", "Cudgel"}
	return fmt.Sprintf("(%d of %s)", id.Number(), seeds[id.Seed()])
}
