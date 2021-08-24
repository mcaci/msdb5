package srvb_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mcaci/msdb5/v2/srvb"
)

func TestCreation(t *testing.T) {
	const url = "localhost:8080/create"
	td := []struct {
		name     string
		method   string
		reqBody  io.Reader
		expected string
	}{
		{"Default game creation", "GET", nil, ""},
		{"Game creation with name", "POST", strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), "newgame"},
		{"Game creation with name and GET", "GET", strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "abc")), "abc"},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, url, tc.reqBody)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			if tc.method == "POST" {
				req.Header.Set("Content-Type", "application/json")
			}
			rec := httptest.NewRecorder()
			srvb.Create(rec, req)
			res := rec.Result()
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.StatusCode)
			}

			var rs struct {
				Name string `json:"name"`
			}
			rserr := json.NewDecoder(res.Body).Decode(&rs)
			if rserr != nil {
				t.Fatalf("%v", rserr)
			}
			if rs.Name != tc.expected {
				t.Fatalf("expecting %v got %v", tc.expected, rs.Name)
			}
		})
	}

}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(srvb.Handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/create", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	expected := "EOF"
	if string(b) != expected {
		t.Fatalf("expecting %v got %v", expected, string(b))
	}
}
