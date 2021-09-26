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
	host    = "localhost:8080"
	newgame = "newgame"
)

func create(b io.Reader) *operation {
	return &operation{url: appendToURL(srvb.CreateURL), hf: srvb.Create, body: b}
}
func join(n string) *operation {
	return &operation{url: appendToURL(srvb.JoinURL), hf: srvb.Join, body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, n, newgame))}
}
func play(b io.Reader) *operation {
	return &operation{url: appendToURL(srvb.PlayURL), hf: srvb.Play, body: b}
}

var withDefaultCreateBody = strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, newgame))

func appendToURL(pattern string) string { return host + pattern }

func TestSrvbOperations(t *testing.T) {
	td := []struct {
		name string
		ops  []*operation
		v    verifier
	}{
		{"Creation with no body gives error", []*operation{
			create(nil),
		}, errWith(http.StatusBadRequest, "")},
		// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
		{"Game creation with name", []*operation{
			create(withDefaultCreateBody),
		}, creationOK(newgame)},
		{"Game creation with error", []*operation{
			create(strings.NewReader(`'{"name":"na"}`)),
		}, errWith(http.StatusInternalServerError, "could not process the request")},
		{"Cannot create two games", []*operation{
			create(withDefaultCreateBody),
			create(strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "errgame"))),
		}, errWith(http.StatusInternalServerError, "one game already created, cannot create more")},
		{"Join with no body gives error", []*operation{
			{body: nil},
		}, errWith(http.StatusBadRequest, "empty request")},
		{"Join with wrong body gives error", []*operation{
			{body: strings.NewReader(`'{"name":"na"}`)},
		}, errWith(http.StatusBadRequest, "could not process the request")},
		{"Join with no create gives error", []*operation{
			join("mary"),
		}, errWith(http.StatusInternalServerError, "not created")},
		{"Join on wrong game", []*operation{
			create(withDefaultCreateBody),
			{body: strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame"))},
		}, errWith(http.StatusInternalServerError, "different name")},
		{"Join with no player name gives error", []*operation{
			create(withDefaultCreateBody),
			{body: strings.NewReader(fmt.Sprintf(`{"game":"%s"}`, newgame))},
		}, errWith(http.StatusInternalServerError, "no player name was given")},
		{"Join with no game name gives error", []*operation{
			create(withDefaultCreateBody),
			{body: strings.NewReader(`{"name":"player"}`)},
		}, errWith(http.StatusInternalServerError, "no game name was given")},
		{"Join with game and player name", []*operation{
			create(withDefaultCreateBody),
			join("mary"),
		}, joinOK("1")},
		{"Two players join", []*operation{
			create(withDefaultCreateBody),
			join("mary"),
			join("michi"),
		}, joinOK("2")},
		{"Three players joining gives error", []*operation{
			create(withDefaultCreateBody),
			join("mary"),
			join("michi"),
			join("onemore"),
		}, errWith(http.StatusInternalServerError, "max players reached")},
		{"Play card with no body gives error", []*operation{
			play(nil),
		}, errWith(http.StatusBadRequest, "empty request")},
		{"Play card with no game gives error", []*operation{
			play(strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", newgame, 1))),
		}, errWith(http.StatusInternalServerError, "not created")},
		{"First player plays ok", []*operation{
			create(withDefaultCreateBody),
			join("mary"),
			join("michi"),
			play(strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", newgame, 1))),
		}, playOK("ok")},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			var res *http.Response
			for _, op := range tc.ops {
				r, err := send(op)
				if err != nil {
					t.Errorf("could perform the operation: %v", err)
				}
				res = r
			}
			if err := verify(res, tc.v); err != nil {
				t.Errorf("test failed because: %v", err)
			}
			srvb.Cleanup(httptest.NewRecorder(), nil)
		})
	}
}
