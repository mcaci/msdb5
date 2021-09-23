package srvb_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/mcaci/msdb5/v2/srvb"
)

const (
	host = "localhost:8080"
)

func url(pattern string) string { return host + pattern }

var (
	createReq = req{urlStr: url(srvb.CreateURL), hf: srvb.Create}
	joinReq   = req{urlStr: url(srvb.JoinURL), hf: srvb.Join}
	playReq   = req{urlStr: url(srvb.PlayURL), hf: srvb.Play}
)

type req struct {
	urlStr string
	hf     http.HandlerFunc
}

func (r req) url() string { return r.urlStr }
func (r req) send(req *http.Request, err error) (*http.Response, error) {
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.hf(rec, req)
	return rec.Result(), nil
}
