package srvp_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/mcaci/msdb5/v2/srvp"
)

const (
	playerNameTest = "tester"
	gameNameTest   = "newgame"
	expectedURL    = "http://localhost:8080/play"
)

type testInfo []uint8

func (testInfo) Name() string      { return gameNameTest }
func (ti testInfo) Cards() []uint8 { return ti }

func TestReactionPlay(t *testing.T) {
	signals := make(chan interface {
		Name() string
		Cards() []uint8
	})
	go func() { signals <- testInfo{1, 2, 3} }()
	actual := srvp.Signal(signals)
	expected := expect(1, 2, 3, 4, 5, 6)
	if !contains(actual, expected) {
		t.Errorf("expecting %q to be inside %q", actual, expected)
	}
}
func TestReactionPlayNotPresent(t *testing.T) {
	signals := make(chan interface {
		Name() string
		Cards() []uint8
	})
	go func() { signals <- testInfo{3, 11, 25} }()
	actual := srvp.Signal(signals)
	expected := expect(1, 5, 20, 37)
	if contains(actual, expected) {
		t.Errorf("expecting %q to be inside %q", actual, expected)
	}
}

func contains(actual interface{}, expected []interface{}) bool {
	for i := range expected {
		if actual != expected[i] {
			log.Println(actual, expected[i])
			continue
		}
		return true
	}
	return false
}

func expect(c ...uint8) []interface{} {
	expected := make([]interface{}, len(c))
	for i := range c {
		expected[i] = struct {
			URL      string
			JsonBody string
		}{URL: expectedURL, JsonBody: fmt.Sprintf(`{"name":"%s","game":"%s","card":"%d"}`, playerNameTest, gameNameTest, c[i])}
	}
	return expected
}
