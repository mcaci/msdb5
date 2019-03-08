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
