package briscola

import (
	"log"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

func IndexOfWinningCard(cardsOnTheTable set.Cards, briscola card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if winningCard(base, other, briscola) == other {
			base = other
			max = i
		}
	}
	return uint8(max)
}

func Winner(cardsOnTheTable []*card.Item, briscola card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if winningCard(*base, *other, briscola) == *other {
			log.Println(*base, *other, briscola)
			base = other
			max = i
		}
	}
	return uint8(max)
}

func winningCard(base, other card.Item, briscola card.Seed) card.Item {
	if &base == nil || doesOtherCardWin(base, other, briscola) {
		base = other
	}
	return base
}

func doesOtherCardWin(first, other card.Item, briscola card.Seed) bool {
	otherIsBriscola := other.Seed() == briscola
	isSameSeed := first.Seed() == other.Seed()
	return (!isSameSeed && otherIsBriscola) || isOtherHigher(first, other)
}

func isOtherHigher(first, other card.Item) bool {
	isSameSeed := first.Seed() == other.Seed()
	isOtherGreaterOnPoints := Points(first) < Points(other)
	isSamePoints := Points(first) == Points(other)
	isOtherGreaterOnNumber := first.Number() < other.Number()
	return isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints)
}
