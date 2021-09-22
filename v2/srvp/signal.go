package srvp

import (
	"fmt"
)

type Carder interface {
	Name() string
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
	}{URL: "http://localhost:8080/play", JsonBody: fmt.Sprintf(`{"name":"%s","game":"%s","card":"%d"}`, "tester", "newgame", sig.Cards()[cardSelF()])}
}
