package round

import "testing"
import "msdb5/card"

func TestWinnerIsAceOfCoin(t *testing.T) {
	first := card.ByID(0)  // 1 Coin
	second := card.ByID(1) // 2 Coin
	third := card.ByID(2)  // 3 Coin
	fourth := card.ByID(3) // 4 Coin
	fifth := card.ByID(4)  // 5 Coin
	briscola := card.Coin

	i := declareWinner(first, second, third, fourth, fifth, briscola)
	if i != 0 {
		t.Fatal("Unexpected winner")
	}
}

func declareWinner(first, second, third, fourth, fifth *(card.Card), briscola card.Seed) uint8 {
	return 0
}
