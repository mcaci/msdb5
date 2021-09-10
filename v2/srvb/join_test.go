package srvb_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/mcaci/msdb5/v2/srvb"
)

const join = joinReq(host + srvb.JoinURL)

type joinReq string

func (j joinReq) url() string { return string(j) }
func (joinReq) send(req *http.Request, err error) (*http.Response, error) {
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	srvb.Join(rec, req)
	return rec.Result(), nil
}

func joinDec(res *http.Response) (string, error) {
	var rs struct {
		Number string `json:"number"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		return "", err
	}
	return rs.Number, nil
}
