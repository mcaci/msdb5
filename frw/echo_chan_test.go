package frw

import (
	"testing"
)

func TestEcho5(t *testing.T) {
	clientChanMap := make(map[int]chan string)
	for i := 0; i < 5; i++ {
		clientChanMap[i] = make(chan string)
	}
	clientID := 1
	client := clientChanMap[clientID]
	go echo(client)
	client <- "Hello"
	if a := <-client; a != "HelloHello" {
		t.Fatalf("Found %s", a)
	}
}

func TestEcho(t *testing.T) {
	clientChan := make(chan string)
	go echo(clientChan)
	clientChan <- "Hello"
	if a := <-clientChan; a != "HelloHello" {
		t.Fatalf("Found %s", a)
	}
}

func echo(client chan string) {
	a := <-client
	// processing
	client <- a + a
}
