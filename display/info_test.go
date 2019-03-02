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
