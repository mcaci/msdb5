package notify

import (
	"errors"
	"testing"

	"golang.org/x/text/language"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type fakeSender struct{}

func (s fakeSender) Sender(string) *player.Player { return player.New() }
func (s fakeSender) Lang() language.Tag           { return language.English }

type fakeRq struct{}

func (r fakeRq) From() string   { return "127.0.0.1" }
func (r fakeRq) Action() string { return "Card#1#Sword" }

type fakeWriter string

func (w *fakeWriter) Write(p []byte) (int, error) {
	str := fakeWriter(string(p))
	*w = str
	return len(str), nil
}

func TestConsoleMsg(t *testing.T) {
	s := new(fakeWriter)
	ToConsole(s, fakeSender{}, fakeRq{})
	if len(*s) == 0 {
		t.Fatalf("Expecting %s but got %s", "", *s)
	}
}

func TestErrConsoleMsg(t *testing.T) {
	s := new(fakeWriter)
	ErrToConsole(s, errors.New("fake err"), fakeSender{}, fakeRq{})
	if len(*s) == 0 {
		t.Fatalf("Expecting %s but got %s", "", *s)
	}
}
