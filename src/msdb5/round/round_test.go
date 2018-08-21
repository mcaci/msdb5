package round

import "testing"
import "msdb5/card"

func TestScenario1WithAceOfCoinWinning(t *testing.T) {
	first, err1 := card.ByID(1)  // 1 Coin
	second, err2 := card.ByID(2) // 2 Coin
	third, err3 := card.ByID(3)  // 3 Coin
	fourth, err4 := card.ByID(4) // 4 Coin
	fifth, err5 := card.ByID(5)  // 5 Coin
	briscola := card.Coin

	if err1 != nil {
		t.Fatal("err1 raised")
	} else if err2 != nil {
		t.Fatal("err2 raised")
	} else if err3 != nil {
		t.Fatal("err3 raised")
	} else if err4 != nil {
		t.Fatal("err4 raised")
	} else if err5 != nil {
		t.Fatal("err5 raised")
	} else {
		i := declareWinner(first, second, third, fourth, fifth, briscola)
		if i != 0 {
			t.Fatal("Unexpected winner")
		}
	}
}

func TestScenario1WithThreeOfCoinWinning(t *testing.T) {
	first, err1 := card.ByID(2)  // 2 Coin
	second, err2 := card.ByID(3) // 3 Coin
	third, err3 := card.ByID(4)  // 4 Coin
	fourth, err4 := card.ByID(5) // 5 Coin
	fifth, err5 := card.ByID(6)  // 6 Coin
	briscola := card.Coin

	if err1 != nil {
		t.Fatal("err1 raised")
	} else if err2 != nil {
		t.Fatal("err2 raised")
	} else if err3 != nil {
		t.Fatal("err3 raised")
	} else if err4 != nil {
		t.Fatal("err4 raised")
	} else if err5 != nil {
		t.Fatal("err5 raised")
	} else {
		i := declareWinner(first, second, third, fourth, fifth, briscola)
		if i != 1 {
			t.Fatalf("Unexpected winner: winner was %d", i)
		}
	}
}
