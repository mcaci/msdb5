package auction

import "testing"

func TestCardsWith80(t *testing.T) {
	number := SideCards(80)
	if number != 0 {
		t.Fatalf("Expected 0 but found %d cards to display", number)
	}
}

func TestCardsWith90(t *testing.T) {
	if number := SideCards(90); number != 1 {
		t.Fatalf("Expected 1 but found %d cards to display", number)
	}
}

func TestCardsWith93(t *testing.T) {
	if number := SideCards(93); number != 1 {
		t.Fatalf("Expected 1 but found %d cards to display", number)
	}
}

func TestCardsWith99(t *testing.T) {
	if number := SideCards(99); number != 1 {
		t.Fatalf("Expected 1 but found %d cards to display", number)
	}
}

func TestCardsWith100(t *testing.T) {
	if number := SideCards(100); number != 2 {
		t.Fatalf("Expected 2 but found %d cards to display", number)
	}
}

func TestCardsWith101(t *testing.T) {
	if number := SideCards(101); number != 2 {
		t.Fatalf("Expected 2 but found %d cards to display", number)
	}
}

func TestCardsWith110(t *testing.T) {
	if number := SideCards(110); number != 3 {
		t.Fatalf("Expected 3 but found %d cards to display", number)
	}
}

func TestCardsWith119(t *testing.T) {
	if number := SideCards(119); number != 3 {
		t.Fatalf("Expected 3 but found %d cards to display", number)
	}
}

func TestCardsWith120(t *testing.T) {
	if number := SideCards(120); number != 5 {
		t.Fatalf("Expected 5 but found %d cards to display", number)
	}
}
