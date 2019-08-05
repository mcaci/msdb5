package player

import (
	"fmt"
	"testing"
)

func initTest() *Player {
	p := New()
	p.RegisterAs("Michi")
	p.Join("127.0.0.1")
	p.Attach(make(chan []byte, 500))
	return p
}

func TestJoinPlayerName(t *testing.T) {
	if p := initTest(); p.Name() != "Michi" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerNameNotEmpty(t *testing.T) {
	if p := initTest(); IsNameEmpty(p) {
		t.Fatal("Unexpected name being empty")
	}
}

func TestJoinPlayerHost(t *testing.T) {
	if p := initTest(); !p.IsSameHost("127.0.0.1") {
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
	go fmt.Fprintf(p, "Hello")
	if str := <-p.info; string(str) != "Hello" {
		t.Fatal("Unexpected message: ", str)
	}
}

func TestMultpleWriteToPlayer(t *testing.T) {
	pls := []*Player{initTest(), initTest(), initTest()}
	for _, p := range pls {
		go func(pl *Player) {
			fmt.Fprint(pl, "Hello")
			fmt.Fprint(pl, "It's me")
			fmt.Fprintln(pl, "Mario")
		}(p)
	}
	for _, p := range pls {
		if str := <-p.info; string(str) != "HelloIt's meMario" {
			t.Fatalf("Unexpected message: %s", str)
		}
	}
}
