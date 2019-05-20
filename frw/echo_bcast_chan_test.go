package frw

import (
	"testing"
)

func TestBroadcast5(t *testing.T) {
	clientChanMap := make(map[int]chan string)
	for i := 0; i < 5; i++ {
		clientChanMap[i] = make(chan string)
	}
	clientID := 1
	client := clientChanMap[clientID]
	go broadcast(clientChanMap)
	client <- "Hello"

	for range clientChanMap {
		select {
		case <-clientChanMap[0]:
		case <-clientChanMap[1]:
		case <-clientChanMap[2]:
		case <-clientChanMap[3]:
		case a := <-clientChanMap[4]:
			if a != "HelloHello" {
				t.Fatalf("Found %s", a)
			}
		}
	}
}

func broadcast(clients map[int]chan string) {
	var a string
	for _, client := range clients {
		select {
		case a = <-client:
		default:
			continue
		}
	}
	// processing
	a = a + a
	for _, client := range clients {
		client <- a
	}
}
