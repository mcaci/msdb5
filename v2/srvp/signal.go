package srvp

import "fmt"

func Signal(signals <-chan struct {
	Name    string
	CardIDs []uint8
}) struct {
	URL      string
	JsonBody string
} {
	<-signals
	return struct {
		URL      string
		JsonBody string
	}{URL: "http://localhost:8080/play", JsonBody: fmt.Sprintf(`{"name":"%s","game":"%s","card":"%d"}`, "tester", "newgame", 1)}
}
