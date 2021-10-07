package srvb

import "net/http"

func Handler() http.Handler {
	mux := http.NewServeMux()
	g := Game{}
	mux.HandleFunc(CreateURL, g.Create)
	mux.HandleFunc(JoinURL, g.Join)
	mux.HandleFunc(PlayURL, g.Play)
	return mux
}
