package srvb

import "net/http"

func Handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/create", Create)
	return r
}
