package srvb_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/mcaci/msdb5/v2/srvb"
)

const (
	host   = "localhost:8080"
	create = createReq(host + srvb.CreateURL)
	join   = joinReq(host + srvb.JoinURL)
	status = statusReq(host + srvb.StatusURL)
)

type createReq string

func (c createReq) url() string { return string(c) }
func (createReq) send(req *http.Request, err error) (*http.Response, error) {
	return send(req, err, srvb.Create)
}

type joinReq string

func (j joinReq) url() string { return string(j) }
func (joinReq) send(req *http.Request, err error) (*http.Response, error) {
	return send(req, err, srvb.Join)
}

func send(req *http.Request, err error, hf http.HandlerFunc) (*http.Response, error) {
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	hf(rec, req)
	return rec.Result(), nil
}

type statusReq string

func (s statusReq) url() string { return string(s) }
func (statusReq) send(req *http.Request, err error) (*http.Response, error) {
	return send(req, err, srvb.Status)
}
