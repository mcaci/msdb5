package card

// ID is the id of a card from 1 to 40
type ID uint8

// HCard is a card represented by string
type HCard struct {
	number, seed string
}

// Creator interface to represent what's needed to create a card
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

// ByName func
func ByName(number, seed string) (Card, error) {
	hc := HCard{number, seed}
	var c Card
	var err error
	if c.number, err = hc.toNumber(); err == nil {
		c.seed, err = hc.toSeed()
	}
	return c, err
}
