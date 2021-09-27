package srvp

import (
	"fmt"
)

const playURL = "http://localhost:8080/play"

type Carder interface {
	Name() string
	Game() string
	Cards() []uint8
}

func Signal(signals <-chan Carder, cardSelF func() int) struct {
	URL      string
	JsonBody string
} {
	sig := <-signals
	return struct {
		URL      string
		JsonBody string
	}{URL: playURL, JsonBody: fmt.Sprintf(`{"name":"%s","game":"%s","card":"%d"}`, sig.Name(), sig.Game(), sig.Cards()[cardSelF()])}
}
