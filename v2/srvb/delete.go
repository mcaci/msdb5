package srvb

import (
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	g = nil
	fmt.Fprint(w, "Delete done")
}
