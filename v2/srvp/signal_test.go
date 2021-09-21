package srvp_test

import (
	"fmt"
	"testing"

	"github.com/mcaci/msdb5/v2/srvp"
)

func TestReactionPlay(t *testing.T) {
	signals := make(chan struct {
		Name    string
		CardIDs []uint8
	})
	go func() {
		signals <- struct {
			Name    string
			CardIDs []uint8
		}{
			Name:    "testgame",
			CardIDs: []uint8{1, 2, 3},
		}
	}()
	actual := srvp.Signal(signals)
	expected := struct {
		URL      string
		JsonBody string
	}{URL: "http://localhost:8080/play", JsonBody: fmt.Sprintf(`{"name":"%s","game":"%s","card":"%d"}`, "tester", "newgame", 1)}
	if actual != expected {
		t.Errorf("expecting %q to be %q", expected, actual)
	}
}
