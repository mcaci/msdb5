package srvb_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mcaci/msdb5/v2/srvb"
)

func TestSrvbOperations(t *testing.T) {
	td := []struct {
		name               string
		setups             []setup
		expectedStatusCode int
		expectedMsg        string
		decoder            func(res *http.Response) (string, error)
	}{
		{"Creation with no body gives error", []setup{{body: nil, r: create}}, http.StatusBadRequest, "", koDec},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", []setup{{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create}}, http.StatusOK, "newgame", creationDec},
		{"Game creation with error", []setup{{body: strings.NewReader(`'{"name":"na"}`), r: create}}, http.StatusInternalServerError, "could not process the request", koDec},
		{"Cannot create two games", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "gg")), r: create},
			{body: strings.NewReader(`{"name":"errgame"}`), r: create},
		}, http.StatusInternalServerError, "one game already created, cannot create more", koDec},
		{"Join with no body gives error", []setup{{body: nil, r: join}}, http.StatusBadRequest, "empty request", koDec},
		{"Join with wrong body gives error", []setup{{body: strings.NewReader(`'{"name":"na"}`), r: join}}, http.StatusBadRequest, "could not process the request", koDec},
		{"Join with no create gives error", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), r: join},
		}, http.StatusInternalServerError, "not created", koDec},
		{"Join on wrong game", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame")), r: join},
		}, http.StatusInternalServerError, "different name", koDec},
		{"Join with no player name gives error", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(`{"game":"newgame"}`), r: join},
		}, http.StatusInternalServerError, "no player name was given", koDec},
		{"Join with no game name gives error", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(`{"name":"player"}`), r: join},
		}, http.StatusInternalServerError, "no game name was given", koDec},
		{"Join with game and player name", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), r: join},
		}, http.StatusOK, "1", joinDec},
		{"Two players join", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), r: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), r: join},
		}, http.StatusOK, "2", joinDec},
		{"Three players joining gives error", []setup{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), r: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), r: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), r: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "onemore", "newgame")), r: join},
		}, http.StatusInternalServerError, "max players reached", koDec},
	}
	for _, tc := range td {
		tester := testResWith(tc.expectedStatusCode, tc.decoder)
		t.Run(tc.name, func(t *testing.T) {
			res, err := exec(tc.setups...)
			if err != nil {
				t.Errorf("could perform the setup: %v", err)
			}
			defer res.Body.Close()
			if err := tester(res, tc.expectedMsg); err != nil {
				t.Errorf("test failed because: %v", err)
			}
			srvb.Cleanup(httptest.NewRecorder(), nil)
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
	for _, el := range s {
		r, err := el.r.send(http.NewRequest(http.MethodPost, el.r.url(), el.body))
		if err != nil {
			return nil, err
		}
		res = r
	}
	return res, nil
}
