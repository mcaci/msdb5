package phase

import (
	"testing"
)

type fakeGamePhase ID

func (p fakeGamePhase) Phase() ID { return ID(p) }

type fakeAction string

func (rq fakeAction) Action() string { return string(rq) }

func TestVerifyPhaseWithNoErr(t *testing.T) {
	err := Check(fakeGamePhase(0), fakeAction("Join"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPhaseWithErr(t *testing.T) {
	err := Check(fakeGamePhase(4), fakeAction("Join"))
	if err == nil {
		t.Fatal("Error was expected")
	}
}
