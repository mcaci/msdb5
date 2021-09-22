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
		name string
		ops  []operation
		v    verifier
	}{
		{"Creation with no body gives error", []operation{{body: nil, req: create}}, errWith(http.StatusBadRequest, "")},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", []operation{{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create}}, creationOK("newgame")},
		{"Game creation with error", []operation{{body: strings.NewReader(`'{"name":"na"}`), req: create}}, errWith(http.StatusInternalServerError, "could not process the request")},
		{"Cannot create two games", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "gg")), req: create},
			{body: strings.NewReader(`{"name":"errgame"}`), req: create},
		}, errWith(http.StatusInternalServerError, "one game already created, cannot create more")},
		{"Join with no body gives error", []operation{{body: nil, req: join}}, errWith(http.StatusBadRequest, "empty request")},
		{"Join with wrong body gives error", []operation{{body: strings.NewReader(`'{"name":"na"}`), req: join}}, errWith(http.StatusBadRequest, "could not process the request")},
		{"Join with no create gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
		}, errWith(http.StatusInternalServerError, "not created")},
		{"Join on wrong game", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame")), req: join},
		}, errWith(http.StatusInternalServerError, "different name")},
		{"Join with no player name gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(`{"game":"newgame"}`), req: join},
		}, errWith(http.StatusInternalServerError, "no player name was given")},
		{"Join with no game name gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(`{"name":"player"}`), req: join},
		}, errWith(http.StatusInternalServerError, "no game name was given")},
		{"Join with game and player name", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
		}, joinOK("1")},
		{"Two players join", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), req: join},
		}, joinOK("2")},
		{"Three players joining gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), req: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "onemore", "newgame")), req: join},
		}, errWith(http.StatusInternalServerError, "max players reached")},
		{"Play card with no body gives error", []operation{{body: nil, req: play}}, errWith(http.StatusBadRequest, "empty request")},
		{"Play card with no game gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d"}`, "onemore", "newgame", 1)), req: play},
		}, errWith(http.StatusInternalServerError, "not created")},
		{"First player plays, ok", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame")), req: join},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d"}`, "onemore", "mary", 1)), req: play},
		}, playOK("")},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			var res *http.Response
			for _, op := range tc.ops {
				r, err := op.req.send(http.NewRequest(http.MethodPost, op.req.url(), op.body))
				if err != nil {
					t.Errorf("could perform the operation: %v", err)
				}
				res = r
			}
			defer res.Body.Close()
			if err := verify(res, tc.v); err != nil {
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
