package srvp_test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/mcaci/msdb5/v3/srvp"
)

const (
	playerNameTest = "tester"
	gameNameTest   = "newgame"
	expectedURL    = "http://localhost:8080/play"
)

type testInfo []uint8

func (testInfo) Game() string      { return gameNameTest }
func (testInfo) Name() string      { return playerNameTest }
func (ti testInfo) Cards() []uint8 { return ti }

func TestSignalPlay(t *testing.T) {
	t.Parallel()
	td := []struct {
		name     string
		in       srvp.Carder
		expected []interface{}
		cntF     func(actual interface{}, expected []interface{}) bool
		msg      string
	}{
		{name: "play on signal can select cards from signal input", in: testInfo{1, 2, 3}, expected: listOf(1, 2, 3, 4, 5, 6),
			cntF: notContains, msg: "to be inside"},
		{name: "play on signal cannot select cards not from signal input", in: testInfo{3, 11, 25}, expected: listOf(1, 5, 20, 37),
			cntF: contains, msg: "not to be inside"},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			signals := make(chan srvp.Carder)
			go func() { signals <- tc.in }()
			actual := srvp.Signal(signals, func() int { return rand.Intn(3) })
			if tc.cntF(actual, tc.expected) {
				t.Errorf("expecting %q to be inside %q", actual, tc.expected)
			}
		})
	}
}

func notContains(actual interface{}, expected []interface{}) bool { return !contains(actual, expected) }
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

func listOf(c ...uint8) []interface{} {
	expected := make([]interface{}, len(c))
	for i := range c {
		expected[i] = struct {
			URL      string
			JsonBody string
		}{URL: expectedURL, JsonBody: fmt.Sprintf(`{"name":"%s","game":"%s","card":"%d"}`, playerNameTest, gameNameTest, c[i])}
	}
	return expected
}
