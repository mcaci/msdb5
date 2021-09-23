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

const (
	newgame = "newgame"
)

func TestSrvbOperations(t *testing.T) {
	td := []struct {
		name string
		ops  []operation
		v    verifier
	}{
		{"Creation with no body gives error", []operation{{body: nil, req: createReq}}, errWith(http.StatusBadRequest, "")},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", []operation{create()}, creationOK(newgame)},
		{"Game creation with error", []operation{{body: strings.NewReader(`'{"name":"na"}`), req: createReq}}, errWith(http.StatusInternalServerError, "could not process the request")},
		{"Cannot create two games", []operation{
			create(),
			createN("errgame"),
		}, errWith(http.StatusInternalServerError, "one game already created, cannot create more")},
		{"Join with no body gives error", []operation{{body: nil, req: joinReq}}, errWith(http.StatusBadRequest, "empty request")},
		{"Join with wrong body gives error", []operation{{body: strings.NewReader(`'{"name":"na"}`), req: joinReq}}, errWith(http.StatusBadRequest, "could not process the request")},
		{"Join with no create gives error", []operation{join("mary")}, errWith(http.StatusInternalServerError, "not created")},
		{"Join on wrong game", []operation{
			create(),
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame")), req: joinReq},
		}, errWith(http.StatusInternalServerError, "different name")},
		{"Join with no player name gives error", []operation{
			create(),
			{body: strings.NewReader(fmt.Sprintf(`{"game":"%s"}`, newgame)), req: joinReq},
		}, errWith(http.StatusInternalServerError, "no player name was given")},
		{"Join with no game name gives error", []operation{
			create(),
			{body: strings.NewReader(`{"name":"player"}`), req: joinReq},
		}, errWith(http.StatusInternalServerError, "no game name was given")},
		{"Join with game and player name", []operation{
			create(),
			join("mary"),
		}, joinOK("1")},
		{"Two players join", []operation{
			create(),
			join("mary"),
			join("michi"),
		}, joinOK("2")},
		{"Three players joining gives error", []operation{
			create(),
			join("mary"),
			join("michi"),
			join("onemore"),
		}, errWith(http.StatusInternalServerError, "max players reached")},
		{"Play card with no body gives error", []operation{{body: nil, req: playReq}}, errWith(http.StatusBadRequest, "empty request")},
		{"Play card with no game gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", newgame, 1)), req: playReq},
		}, errWith(http.StatusInternalServerError, "not created")},
		{"First player plays ok", []operation{
			create(),
			join("mary"),
			join("michi"),
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", newgame, 1)), req: playReq},
		}, playOK("ok")},
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

var create = func() operation { return createN(newgame) }
var createN = func(n string) operation {
	return operation{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, n)), req: createReq}
}
var join = func(n string) operation {
	return operation{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, n, newgame)), req: joinReq}
}
