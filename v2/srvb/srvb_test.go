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
		ops                []operation
		expectedStatusCode int
		expectedMsg        string
		decoder            func(res *http.Response) (string, error)
	}{
		{"Creation with no body gives error", []operation{{body: nil, req: create}}, http.StatusBadRequest, "", koDec},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", []operation{{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create}}, http.StatusOK, "newgame", creationDec},
		{"Game creation with error", []operation{{body: strings.NewReader(`'{"name":"na"}`), req: create}}, http.StatusInternalServerError, "could not process the request", koDec},
		{"Cannot create two games", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "gg")), req: create},
			{body: strings.NewReader(`{"name":"errgame"}`), req: create},
		}, http.StatusInternalServerError, "one game already created, cannot create more", koDec},
		{"Join with no body gives error", []operation{{body: nil, req: join}}, http.StatusBadRequest, "empty request", koDec},
		{"Join with wrong body gives error", []operation{{body: strings.NewReader(`'{"name":"na"}`), req: join}}, http.StatusBadRequest, "could not process the request", koDec},
		{"Join with no create gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
		}, http.StatusInternalServerError, "not created", koDec},
		{"Join on wrong game", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame")), req: join},
		}, http.StatusInternalServerError, "different name", koDec},
		{"Join with no player name gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(`{"game":"newgame"}`), req: join},
		}, http.StatusInternalServerError, "no player name was given", koDec},
		{"Join with no game name gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(`{"name":"player"}`), req: join},
		}, http.StatusInternalServerError, "no game name was given", koDec},
		{"Join with game and player name", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
		}, http.StatusOK, "1", joinDec},
		{"Two players join", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), req: join},
		}, http.StatusOK, "2", joinDec},
		{"Three players joining gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), req: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "onemore", "newgame")), req: join},
		}, http.StatusInternalServerError, "max players reached", koDec},
	}
	for _, tc := range td {
		tester := testResWith(tc.expectedStatusCode, tc.decoder)
		t.Run(tc.name, func(t *testing.T) {
			res, err := exec(tc.ops...)
			if err != nil {
				t.Errorf("could perform the operation: %v", err)
			}
			defer res.Body.Close()
			if err := tester(res, tc.expectedMsg); err != nil {
				t.Errorf("test failed because: %v", err)
			}
			srvb.Cleanup(httptest.NewRecorder(), nil)
		})
	}
}

type operation struct {
	body io.Reader
	req  interface {
		url() string
		send(*http.Request, error) (*http.Response, error)
	}
}

func exec(ops ...operation) (*http.Response, error) {
	var res *http.Response
	for _, op := range ops {
		r, err := op.req.send(http.NewRequest(http.MethodPost, op.req.url(), op.body))
		if err != nil {
			return nil, err
		}
		res = r
	}
	return res, nil
}
