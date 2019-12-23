package player

import (
	"fmt"
	"testing"
)

func initTest() *Player {
	p := New()
	p.RegisterAs("Michi")
	p.Join("127.0.0.1")
	p.Attach(make(chan []byte, 3))
	return p
}

func TestNewPlayersAreNotSame(t *testing.T) {
	if Matching(initTest())(initTest()) {
		t.Fatal("Unexpected players being equal")
	}
}

func TestJoinPlayerName(t *testing.T) {
	if p := initTest(); p.Name() != "Michi" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerNameNotEmpty(t *testing.T) {
	if IsNameEmpty(initTest()) {
		t.Fatal("Unexpected name being empty")
	}
}

func TestJoinPlayerHost(t *testing.T) {
	if !MatchingHost("127.0.0.1")(initTest()) {
		t.Fatal("Unexpected host")
	}
}

func TestJoinPlayerPileIsEmpty(t *testing.T) {
	if p := initTest(); len(*p.Pile()) != 0 {
		t.Fatal("Pile should be empty")
	}
}

func TestWriteToPlayer(t *testing.T) {
	p := initTest()
	go fmt.Fprint(p, "Hello")
	if str := <-p.info; string(str) != "Hello" {
		t.Fatal("Unexpected message: ", str)
	}
}
