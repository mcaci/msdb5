package card

// Card type
type Card struct {
	number uint8
	seed   Seed
}

// StrCard is a card represented by string
type StrCard struct {
	number, seed string
}

// ID is the id of a card from 1 to 40
type ID uint8

// Creator interface to represent what's needed to create a card
type Creator interface {
	toNumber() (uint8, error)
	toSeed() (Seed, error)
}

// By func
func By(sCard Creator) (Card, error) {
	var c Card
	var err error
	if c.number, err = sCard.toNumber(); err == nil {
		c.seed, err = sCard.toSeed()
	}
	return c, err
}

// ByName func
func ByName(number, seed string) (Card, error) {
	sCard := StrCard{number, seed}
	var c Card
	var err error
	if c.number, err = sCard.toNumber(); err == nil {
		c.seed, err = sCard.toSeed()
	}
	return c, err
}

// ID func
func (card *Card) ID() ID {
	return ID(card.number + (uint8)(card.seed)*10)
}

// Number func
func (card *Card) Number() uint8 {
	return card.number
}

// Seed func
func (card *Card) Seed() Seed {
	return card.seed
}

// Points func
func (card *Card) Points() uint8 {
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
