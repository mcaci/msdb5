package srvb

import "net/http"

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(CreateURL, Create)
	mux.HandleFunc(JoinURL, Join)
	mux.HandleFunc(PlayURL, Play)
	return mux
}
