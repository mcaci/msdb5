package srvb_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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
