package srvb

import "net/http"

func Handler() http.Handler {
	mux := http.NewServeMux()
	g := Game{}
	mux.HandleFunc(CreateURL, g.Create)
	mux.HandleFunc("/CreateURL", Create)
	mux.HandleFunc(JoinURL, g.Join)
	mux.HandleFunc("/JoinURL", Join)
	mux.HandleFunc(PlayURL, Play)
	return mux
}
