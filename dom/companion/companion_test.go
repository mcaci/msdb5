package companion

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func TestCardNominatedWithOriginInfo(t *testing.T) {
	comp := New(1, player.New())
	if comp == nil {
		t.Fatal("Companion was not created")
	}
}

func TestCardNominatedWithOriginInfoCard(t *testing.T) {
	comp := New(19, player.New())
	if comp.Card() != 19 {
		t.Fatal("Companion card is not 9 of Cups")
	}
}

func TestPlayerNominatedIsNotNil(t *testing.T) {
	comp := New(2, player.New())
	if comp.Ref() == nil {
		t.Fatal("Player is nil")
	}
}
