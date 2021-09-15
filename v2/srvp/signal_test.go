package srvp_test

import (
	"testing"

	"github.com/mcaci/msdb5/v2/srvp"
)

func TestReactionPlay(t *testing.T) {
	signals := make(chan struct{})
	go func() { signals <- struct{}{} }()
	actual := srvp.Signal(signals)
	expected := "play"
	if actual != expected {
		t.Errorf("expecting %q to be %q", expected, actual)
	}
}
