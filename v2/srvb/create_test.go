package srvb_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/mcaci/msdb5/v2/srvb"
)

const create = createReq("localhost:8080/create")

type createReq string

func (c createReq) url() string { return string(c) }
func (createReq) send(req *http.Request, err error) (*http.Response, error) {
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	srvb.Create(rec, req)
	return rec.Result(), nil
}

func creationRes(res *http.Response) (string, error) {
	var rs struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		return "", err
	}
	return rs.Name, nil
}
