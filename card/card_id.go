package card

// ID is the id of a card from 1 to 40
type ID uint8

// HCard is a card represented by string
type HCard struct {
	number, seed string
}

type Creator interface {
	toNumber() (uint8, error)
	toSeed() (Seed, error)
}

// By func
func By(hc Creator) (Card, error) {
	var c Card
	var err error
	if c.number, err = hc.toNumber(); err == nil {
		c.seed, err = hc.toSeed()
	}
	return c, err
}

func ByName(number, seed string) (Card, error) {
	return By(HCard{number, seed})
}

func toZeroBased(id ID) uint8 {
	return uint8(id) - 1
}

func fromZeroBased(index uint8) ID {
	return ID(index + 1)
}
