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

var (
	create = req{urlStr: host + srvb.CreateURL, hf: srvb.Create}
	join   = req{urlStr: host + srvb.JoinURL, hf: srvb.Join}
	status = req{urlStr: host + srvb.StatusURL, hf: srvb.Status}
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
