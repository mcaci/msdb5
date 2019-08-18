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

func initTestWithName(name string) *Player {
	p := New()
	p.RegisterAs(name)
	p.Join("127.0.0.1")
	p.Attach(make(chan []byte))
	return p
}

func TestJoinPlayerName(t *testing.T) {
	if p := initTest(); p.Name() != "Michi" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerNameNotEmpty(t *testing.T) {
	if p := initTest(); p.Name() != "" {
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
	go fmt.Fprint(p, "Hello")
	if str := <-p.info; string(str) != "Hello" {
		t.Fatal("Unexpected message: ", str)
	}
}
