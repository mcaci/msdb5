package briscola_test

import (
	"log"
	"testing"

	"github.com/mcaci/msdb5/v3/briscola"
)

type inTest struct {
	name    string
	cardIdx uint8
}

func (i inTest) Name() string { return i.name }
func (i inTest) Idx() uint8   { return i.cardIdx }

func TestEmptyPlayWithError(t *testing.T) {
	t.Parallel()
	in := inTest{"p1", 0}
	gamestate := briscola.NewGame(briscola.WithDefaultOptions)
	_, err := briscola.Play(gamestate, in)
	if err == nil {
		t.Error("Expecting an error but all went fine")
	}
}

func TestPlayActionOk(t *testing.T) {
	t.Parallel()
	gamestate := briscola.NewGame(briscola.WithDefaultOptions)
	briscola.Register("p1", gamestate)
	briscola.Register("p2", gamestate)
	out, err := briscola.Play(gamestate, inTest{"p1", 0})
	if err != nil {
		t.Errorf("Not expecting error but got %v", err)
	}
	if out.Pl.Name() == "" {
		t.Errorf("Expecting a player but got %v", out.Pl)
	}
	if len(*out.Brd) == 0 {
		t.Error("Expecting the board to be filled but got an empty board")
	}
}

func TestSamePlayerCannotPlayTwice(t *testing.T) {
	t.Parallel()
	gamestate := briscola.NewGame(briscola.WithDefaultOptions)
	briscola.Register("p1", gamestate)
	briscola.Register("p2", gamestate)
	briscola.Play(gamestate, inTest{"p1", 0})
	_, err := briscola.Play(gamestate, inTest{"p1", 0})
	if err == nil {
		t.Error("Expecting an error here but got nil")
	}
}

func TestSecondPlayerPlaysOK(t *testing.T) {
	t.Parallel()
	gamestate := briscola.NewGame(briscola.WithDefaultOptions)
	briscola.Register("p1", gamestate)
	briscola.Register("p2", gamestate)
	briscola.Play(gamestate, inTest{"p1", 0})
	out, err := briscola.Play(gamestate, inTest{"p2", 0})
	if err != nil {
		t.Errorf("Not expecting error but got %v", err)
	}
	if len(*(*gamestate.Players())[0].Hand()) != 3 {
		t.Errorf("Expecting Player 1 to have 3 card but was %d", len(*(*gamestate.Players())[0].Hand()))
	}
	if len(*(*gamestate.Players())[1].Hand()) != 3 {
		t.Errorf("Expecting Player 1 to have 3 card but was %d", len(*(*gamestate.Players())[1].Hand()))
	}
	if out.Pl.Name() == "" {
		t.Errorf("Expecting a player but got %v", out.Pl)
	}
	if len(*out.Brd) != 0 {
		t.Errorf("Expecting the cards on the board to be already collected but got %v", *out.Brd)
	}
}

func TestCannotPlayCardOutsideHand(t *testing.T) {
	t.Parallel()
	gamestate := briscola.NewGame(briscola.WithDefaultOptions)
	briscola.Register("p1", gamestate)
	briscola.Register("p2", gamestate)
	_, err := briscola.Play(gamestate, inTest{"p1", 3})
	if err == nil {
		t.Errorf("Not expecting error but got %v", err)
	}
}

func TestFullGame(t *testing.T) {
	t.Parallel()
	gamestate := briscola.NewGame(briscola.WithDefaultOptions)
	briscola.Register("p1", gamestate)
	briscola.Register("p2", gamestate)
	for i := 0; len(*gamestate.InTurn().Hand()) > 0 && i < 40; i++ {
		log.Println(i, gamestate)
		out, err := briscola.Play(gamestate, inTest{gamestate.InTurn().Name(), 0})
		if err != nil {
			t.Fatalf("Not expecting error but got %v", err)
		}
		if out.Pl.Name() == "" {
			t.Errorf("Expecting a player but got %v", out.Pl)
		}
		if len(*out.Brd) > 2 {
			t.Errorf("Cards on the board cannot be more than two but got %v", *out.Brd)
		}
	}
}
