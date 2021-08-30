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

const url = "localhost:8080/create"

func TestCreation(t *testing.T) {
	td := []struct {
		name     string
		method   string
		reqBody  io.Reader
		expected string
	}{
		{"Default game creation", "GET", nil, ""},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
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
				t.Fatalf("expected status OK; got %v", res.StatusCode)
			}
			actual, err := func(res *http.Response) (string, error) {
				var rs struct {
					Name string `json:"name"`
				}
				err := json.NewDecoder(res.Body).Decode(&rs)
				if err != nil {
					return "", err
				}
				return rs.Name, nil
			}(res)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}
			expected := tc.expected
			if actual != expected {
				t.Fatalf("expecting %v got %v", expected, actual)
			}
			cleanup(t)
		})
	}
}

func TestErrCreation(t *testing.T) {
	sendReq := func(n string) (*http.Response, error) {
		req, err := http.NewRequest("POST", url, strings.NewReader(fmt.Sprintf(`'{"name":"%s"}`, n)))
		if err != nil {
			return nil, fmt.Errorf("could not create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		srvb.Create(rec, req)
		return rec.Result(), nil
	}
	n2 := "errgame"
	res2, err := sendReq(n2)
	if err != nil {
		t.Fatalf("could not send the request: %v", err)
	}
	defer res2.Body.Close()
	if res2.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", res2.StatusCode)
	}
	b, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		t.Fatalf("could not read error response: %v", err)
	}
	expected := "could not process the request"
	if actual := string(b); !strings.Contains(actual, expected) {
		t.Fatalf("expecting %q got %q", expected, actual)
	}
	cleanup(t)
}

func TestCanCreateOnlyOneGame(t *testing.T) {
	sendReq := func(n string) (*http.Response, error) {
		req, err := http.NewRequest("POST", url, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, n)))
		if err != nil {
			return nil, fmt.Errorf("could not create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		srvb.Create(rec, req)
		return rec.Result(), nil
	}
	n1 := "gg"
	res1, err := sendReq(n1)
	if err != nil {
		t.Fatalf("could not send first request: %v", err)
	}
	defer res1.Body.Close()
	if res1.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", res1.StatusCode)
	}
	actual, err := func(res *http.Response) (string, error) {
		var rs struct {
			Name string `json:"name"`
		}
		err := json.NewDecoder(res.Body).Decode(&rs)
		if err != nil {
			return "", err
		}
		return rs.Name, nil
	}(res1)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	if actual != n1 {
		t.Fatalf("expecting %v got %v", n1, actual)
	}
	n2 := "errgame"
	res2, err := sendReq(n2)
	if err != nil {
		t.Fatalf("could not send second request: %v", err)
	}
	defer res2.Body.Close()
	if res2.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected status InternalServerError; got %v", res2.StatusCode)
	}
	b, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		t.Fatalf("could not read error response: %v", err)
	}
	expected := "one game already created, cannot create more\n"
	if actual := string(b); actual != expected {
		t.Fatalf("expecting %q got %q", expected, actual)
	}
	cleanup(t)
}

func cleanup(t *testing.T) {
	const durl = "localhost:8080/dgame"
	dreq, err := http.NewRequest(http.MethodDelete, durl, nil)
	if err != nil {
		t.Fatalf("could not send DELETE request: %v", err)
	}
	srvb.Delete(httptest.NewRecorder(), dreq)
}
