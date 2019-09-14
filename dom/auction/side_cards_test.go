package auction

import "testing"

func sideCardsTest(t *testing.T, score Score, expected uint8) {
	if actual := SideCards(score); actual != expected {
		t.Fatalf("Expected %d but found %d cards to display", expected, actual)
	}
}

func TestCardsWith80(t *testing.T) {
	sideCardsTest(t, 80, 0)
}

func TestCardsWith90(t *testing.T) {
	sideCardsTest(t, 90, 1)
}

func TestCardsWith93(t *testing.T) {
	sideCardsTest(t, 93, 1)
}

func TestCardsWith99(t *testing.T) {
	sideCardsTest(t, 99, 1)
}

func TestCardsWith100(t *testing.T) {
	sideCardsTest(t, 100, 2)
}

func TestCardsWith101(t *testing.T) {
	sideCardsTest(t, 101, 2)
}

func TestCardsWith110(t *testing.T) {
	sideCardsTest(t, 110, 3)
}

func TestCardsWith119(t *testing.T) {
	sideCardsTest(t, 119, 3)
}

func TestCardsWith120(t *testing.T) {
	sideCardsTest(t, 120, 5)
}
