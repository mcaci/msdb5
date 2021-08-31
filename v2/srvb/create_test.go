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

const crUrl = "localhost:8080/create"

func TestCreation(t *testing.T) {
	td := []struct {
		name     string
		method   string
		reqBody  io.Reader
		tester   func(*http.Response, string) error
		expected string
	}{
		{"Default game creation", http.MethodGet, nil, testOK, ""},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", http.MethodPost, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), testOK, "newgame"},
		{"Game creation with name and GET", http.MethodGet, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "abc")), testOK, "abc"},
		{"Game creation with error", http.MethodPost, strings.NewReader(`'{"name":"na"}`), testKO, "could not process the request"},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			res, err := send(http.NewRequest(tc.method, crUrl, tc.reqBody))
			if err != nil {
				t.Fatalf("could not send the request: %v", err)
			}
			defer res.Body.Close()
			if err := tc.tester(res, tc.expected); err != nil {
				t.Fatalf("test failed because: %v", err)
			}
			if err := cleanup(); err != nil {
				t.Fatalf("test passed but cleanup failed because: %v", err)
			}
		})
	}
}

func TestCanCreateOnlyOneGame(t *testing.T) {
	td := []struct {
		name     string
		reqBody  io.Reader
		tester   func(*http.Response, string) error
		expected string
	}{
		{"first request", strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "gg")), testOK, "gg"},
		{"second request", strings.NewReader(`{"name":"errgame"}`), testKO, "one game already created, cannot create more"},
	}
	for _, tc := range td {
		res, err := send(http.NewRequest(http.MethodPost, crUrl, tc.reqBody))
		if err != nil {
			t.Fatalf("could not send the request: %v", err)
		}
		defer res.Body.Close()
		if err := tc.tester(res, tc.expected); err != nil {
			t.Fatalf("test failed because: %v", err)
		}
	}
	if err := cleanup(); err != nil {
		t.Fatalf("test passed but cleanup failed because: %v", err)
	}
}

func send(req *http.Request, err error) (*http.Response, error) {
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	if req.Method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}
	rec := httptest.NewRecorder()
	srvb.Create(rec, req)
	return rec.Result(), nil
}

func testOK(res *http.Response, expected string) error {
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status OK; got %v", res.StatusCode)
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
		return fmt.Errorf("could not read response: %v", err)
	}
	if actual != expected {
		return fmt.Errorf("expecting %v got %v", expected, actual)
	}
	return nil
}

func testKO(res *http.Response, expected string) error {
	if res.StatusCode != http.StatusInternalServerError {
		return fmt.Errorf("expected status InternalServerError; got %v", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not read error response: %v", err)
	}
	if actual := string(b); !strings.Contains(actual, expected) {
		return fmt.Errorf("expecting %q to be in %q", expected, actual)
	}
	return nil
}

func cleanup() error {
	const durl = "localhost:8080/dgame"
	dreq, err := http.NewRequest(http.MethodDelete, durl, nil)
	if err != nil {
		return fmt.Errorf("could not send DELETE request: %v", err)
	}
	srvb.Delete(httptest.NewRecorder(), dreq)
	return nil
}
