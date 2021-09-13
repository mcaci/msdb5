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
		v    interface{ verify(*http.Response) error }
	}{
		{"Creation with no body gives error", []operation{{body: nil, req: create}}, badReqErr("")},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", []operation{{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create}}, crOK("newgame")},
		{"Game creation with error", []operation{{body: strings.NewReader(`'{"name":"na"}`), req: create}}, intSrvErr("could not process the request")},
		{"Cannot create two games", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "gg")), req: create},
			{body: strings.NewReader(`{"name":"errgame"}`), req: create},
		}, intSrvErr("one game already created, cannot create more")},
		{"Join with no body gives error", []operation{{body: nil, req: join}}, badReqErr("empty request")},
		{"Join with wrong body gives error", []operation{{body: strings.NewReader(`'{"name":"na"}`), req: join}}, badReqErr("could not process the request")},
		{"Join with no create gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame")), req: join},
		}, intSrvErr("not created")},
		{"Join on wrong game", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame")), req: join},
		}, intSrvErr("different name")},
		{"Join with no player name gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(`{"game":"newgame"}`), req: join},
		}, intSrvErr("no player name was given")},
		{"Join with no game name gives error", []operation{
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")), req: create},
			{body: strings.NewReader(`{"name":"player"}`), req: join},
		}, intSrvErr("no game name was given")},
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
		}, intSrvErr("max players reached")},
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
			if err := tc.v.verify(res); err != nil {
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
