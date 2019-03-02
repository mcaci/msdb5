package display

import "testing"

func TestPrintOfEmptyDataIsEmpty(t *testing.T) {
	info := NewInfo("", "", "", "")
	if info.Display() != "" {
		t.Fatal("nothing should be in the output")
	}
}

func TestAddAllOfEmptyDataIsEmpty(t *testing.T) {
	info1 := NewInfo("", "", "", "")
	info2 := NewInfo("", "", "", "")
	if Wrap("", info1, info2) == nil {
		t.Fatal("unexpected output")
	}
}

func TestPrintAllOfEmptyDataIsEmpty(t *testing.T) {
	info1 := NewInfo("", "", "", "")
	info2 := NewInfo("", "", "", "")
	infos := Wrap("", info1, info2)
	if All(infos...) != "()" {
		t.Fatal("unexpected output")
	}
}

// type mockStringer uint8

// func (m mockStringer) String() string {
// 	return "1"
// }

// func TestToStringOfMockData(t *testing.T) {
// 	var info1 mockStringer
// 	var info2 mockStringer
// 	if ToString(info1, info2) != "1 1 " {
// 		t.Fatal("unexpected output")
// 	}
// }

// type mockStringers []*mockStringer

// func TestToStringOfTypedMockData(t *testing.T) {
// 	var info1 mockStringer
// 	var info2 mockStringer
// 	infos := mockStringers{&info1, &info2}
// 	if ToString((*infos)...) != "1 1 " {
// 		t.Fatal("unexpected output")
// 	}
// }
