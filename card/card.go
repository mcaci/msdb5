package card

// Data type
type Data struct {
	number uint8
	seed   Seed
}

// StrData is a card represented by string
type StrData struct {
	number, seed string
}

// ID is the id of a card from 1 to 40
type ID uint8

// Creator interface to represent what's needed to create a card
type Creator interface {
	ToNumber() uint8
	ToSeed() Seed
}

// By func
func By(sCard ID) (Data, error) {
	var c Data
	var err error
	if c.number, err = sCard.toNumber(); err == nil {
		c.seed, err = sCard.toSeed()
	}
	return c, err
}

// ByName func
func ByName(number, seed string) (Data, error) {
	sCard := StrData{number, seed}
	var c Data
	var err error
	if c.number, err = sCard.toNumber(); err == nil {
		c.seed, err = sCard.toSeed()
	}
	return c, err
}

// ID func
func (card *Data) ID() ID {
	return ID(card.number + (uint8)(card.seed)*10)
}

// Number func
func (card *Data) Number() uint8 {
	return card.number
}

// Seed func
func (card *Data) Seed() Seed {
	return card.seed
}

// Points func
func (card *Data) Points() uint8 {
	switch card.number {
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
