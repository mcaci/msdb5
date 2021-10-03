package srvb_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/mcaci/msdb5/v3/srvb"
)

type operation struct {
	body io.Reader
	url  string
	hf   http.HandlerFunc
}

func send(op *operation) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, op.url, op.body)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	op.hf(rec, req)
	return rec.Result(), nil
}

func create(b io.Reader) *operation {
	return &operation{url: appendToURL(srvb.CreateURL), hf: srvb.Create, body: b}
}
func join(b io.Reader) *operation {
	return &operation{url: appendToURL(srvb.JoinURL), hf: srvb.Join, body: b}
}
func play(b io.Reader) *operation {
	return &operation{url: appendToURL(srvb.PlayURL), hf: srvb.Play, body: b}
}
func appendToURL(pattern string) string { const host = "localhost:8080"; return host + pattern }
