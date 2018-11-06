package card

func (id ID) toNumber() (uint8, error) {
	return id.toZeroBased()%10 + 1, nil
}

func (id ID) toSeed() (Seed, error) {
	return Seed(id.toZeroBased() / 10), nil
}

// ToNumber func
func (id ID) ToNumber() uint8 {
	n, _ := id.toNumber()
	return n
}

// ToSeed func
func (id ID) ToSeed() Seed {
	s, _ := id.toSeed()
	return s
}

func (id ID) toZeroBased() uint8 {
	return uint8(id) - 1
}
