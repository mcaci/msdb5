package card

import "strconv"

// Seed type
type Seed uint8

const (
	// Coin elements
	Coin Seed = iota
	// Cup elements
	Cup
	// Sword elements
	Sword
	// Cudgel elements
	Cudgel
)

const _SeedName = "CoinCupSwordCudgel"

var _SeedIndex = [...]uint8{0, 4, 7, 12, 18}

func (i Seed) String() string {
	if i >= Seed(len(_SeedIndex)-1) {
		return "Seed(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SeedName[_SeedIndex[i]:_SeedIndex[i+1]]
}
