package srvb

import (
	"fmt"
	"net/http"
)

func Cleanup(w http.ResponseWriter, r *http.Request) {
	g = nil
	fmt.Fprint(w, "Cleanup done")
}
