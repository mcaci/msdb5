package srvb

import "net/http"

func Handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc(CreateURL, Create)
	r.HandleFunc(JoinURL, Join)
	return r
}
