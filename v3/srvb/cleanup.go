package srvb

import (
	"fmt"
	"net/http"
)

const CleanupURL = "/cln"

func Cleanup(w http.ResponseWriter, r *http.Request) {
	g = nil
	fmt.Fprint(w, "Cleanup done")
}
