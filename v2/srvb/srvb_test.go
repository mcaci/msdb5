package srvb_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mcaci/msdb5/v2/srvb"
)

func TestSrvbOperations(t *testing.T) {
	td := []struct {
		name     string
		setups   []setup
		tester   func(*http.Response, string) error
		expected string
	}{
		{"Creation with no body gives error", []setup{{body: nil, r: create}}, testKOWith(http.StatusInternalServerError), ""},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", []setup{{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create}}, testOKFor(creationRes), "newgame"},
		{"Game creation with error", []setup{{body: strings.NewReader(`'{"name":"na"}`), r: create}}, testKOWith(http.StatusInternalServerError), "could not process the request"},
		{"Cannot create two games", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "gg")), r: create},
			{body: strings.NewReader(`{"name":"errgame"}`), r: create},
		}, testKOWith(http.StatusInternalServerError), "one game already created, cannot create more"},
		{"Join with no body gives error", []setup{{body: nil, r: join}}, testKOWith(http.StatusBadRequest), "empty request"},
		{"Join with no create gives error", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), r: join},
		}, testKOWith(http.StatusInternalServerError), "not created"},
		{"Join on wrong game", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame")), r: join},
		}, testKOWith(http.StatusInternalServerError), "different name"},
		{"Join with game and player name", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), r: join},
		}, testOKFor(joinRes), "1"},
		{"Two players join", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), r: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), r: join},
		}, testOKFor(joinRes), "2"},
		{"Three players joining gives error", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), r: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), r: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "onemore", "newgame")), r: join},
		}, testKOWith(http.StatusInternalServerError), "max players reached"},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			res, err := exec(tc.setups...)
			if err != nil {
				t.Errorf("could perform the setup: %v", err)
			}
			defer res.Body.Close()
			if err := tc.tester(res, tc.expected); err != nil {
				t.Errorf("test failed because: %v", err)
			}
			if err := cleanup(); err != nil {
				t.Errorf("test passed but cleanup failed because: %v", err)
			}
		})
	}
}

type setup struct {
	body io.Reader
	r    interface {
		url() string
		send(*http.Request, error) (*http.Response, error)
	}
}

func exec(s ...setup) (*http.Response, error) {
	var res *http.Response
	for _, e := range s {
		r, err := e.r.send(http.NewRequest(http.MethodPost, e.r.url(), e.body))
		if err != nil {
			return nil, err
		}
		res = r
	}
	return res, nil
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

func testOKFor(decode func(res *http.Response) (string, error)) func(*http.Response, string) error {
	return func(res *http.Response, expected string) error {
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("expected status OK; got %v", res.StatusCode)
		}
		actual, err := decode(res)
		if err != nil {
			return fmt.Errorf("could not read response: %v", err)
		}
		if actual != expected {
			return fmt.Errorf("expecting %v got %v", expected, actual)
		}
		return nil
	}
}

func testKOWith(httpStatusCode int) func(*http.Response, string) error {
	return func(res *http.Response, expected string) error {
		if res.StatusCode != httpStatusCode {
			return fmt.Errorf("expected status %d; got %d", httpStatusCode, res.StatusCode)
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
}
