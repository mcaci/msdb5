package briscola_test

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v3/briscola"
)

func TestNewPlayedCards(t *testing.T) {
	t.Parallel()
	pc := briscola.NewPlayedCards(2)
	if len(*pc.Pile()) != 0 {
		t.Errorf("error: expecting 0 but was %d", len(*pc.Pile()))
	}
}

func TestEmptyPlayedCards(t *testing.T) {
	t.Parallel()
	pc := briscola.NewPlayedCards(2)
	pc.Add(*card.MustID(1))
	if len(*pc.Pile()) != 0 {
		t.Errorf("error: expecting 0 but was %d", len(*pc.Pile()))
	}
}

func TestPlayedCards(t *testing.T) {
	t.Parallel()
	pc := briscola.NewPlayedCards(2)
	pc.Add(*card.MustID(1), *card.MustID(2))
	if len(*pc.Pile()) != 2 {
		t.Errorf("error: expecting 2 but was %d", len(*pc.Pile()))
	}
}
