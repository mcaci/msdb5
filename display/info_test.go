package display

import "testing"

func TestPrintOfEmptyDataIsEmpty(t *testing.T) {
	info := NewInfo("", "", "", "")
	if info.PrintIt() != "" {
		t.Fatal("nothing should be in the output")
	}
}

func TestPrintAllOfEmptyDataIsEmpty(t *testing.T) {
	info1 := NewInfo("", "", "", "")
	info2 := NewInfo("", "", "", "")
	if PrintAll(info1, info2) != "" {
		t.Fatal("nothing should be in the output")
	}
}

type mockStringer uint8

func (m mockStringer) String() string {
	return "1"
}

func TestToStringOfMockData(t *testing.T) {
	var info1 mockStringer
	var info2 mockStringer
	if ToString(info1, info2) != "1 1 " {
		t.Fatal("nothing should be in the output")
	}
}
