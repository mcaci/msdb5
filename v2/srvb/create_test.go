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

type crReq struct {
	body     io.Reader
	tester   func(*http.Response, string) error
	expected string
}

func TestCreation(t *testing.T) {
	td := []struct {
		name   string
		method string
		reqs   []crReq
	}{
		{"Default game creation", http.MethodGet, []crReq{{nil, testCreateOK, ""}}},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", http.MethodPost, []crReq{{strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), testCreateOK, "newgame"}}},
		{"Game creation with name and GET", http.MethodGet, []crReq{{strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "abc")), testCreateOK, "abc"}}},
		{"Game creation with error", http.MethodPost, []crReq{{strings.NewReader(`'{"name":"na"}`), testCreateKO, "could not process the request"}}},
		{"Cannot create two games", http.MethodPost, []crReq{
			{strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "gg")), testCreateOK, "gg"},
			{strings.NewReader(`{"name":"errgame"}`), testCreateKO, "one game already created, cannot create more"},
		}},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			for _, req := range tc.reqs {
				res, err := sendCreate(http.NewRequest(tc.method, crUrl, req.body))
				if err != nil {
					t.Fatalf("could not send the request: %v", err)
				}
				defer res.Body.Close()
				if err := req.tester(res, req.expected); err != nil {
					t.Fatalf("test failed because: %v", err)
				}
			}
			if err := cleanup(); err != nil {
				t.Fatalf("test passed but cleanup failed because: %v", err)
			}
		})
	}
}

func sendCreate(req *http.Request, err error) (*http.Response, error) {
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	srvb.Create(rec, req)
	return rec.Result(), nil
}

func testCreateOK(res *http.Response, expected string) error {
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

func testCreateKO(res *http.Response, expected string) error {
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
