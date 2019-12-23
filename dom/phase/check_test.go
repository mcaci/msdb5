package phase

import (
	"testing"
)

type fakeGamePhase struct {
	id  ID
	act string
}

func (p fakeGamePhase) Phase() ID      { return p.id }
func (p fakeGamePhase) Action() string { return p.act }

func TestVerifyPhaseWithNoErr(t *testing.T) {
	err := Check(fakeGamePhase{0, "Join"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPhaseWithErr(t *testing.T) {
	err := Check(fakeGamePhase{4, "Join"})
	if err == nil {
		t.Fatal("Error was expected")
	}
}
