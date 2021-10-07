package srvb_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mcaci/msdb5/v3/srvb"
)

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(srvb.Handler())
	defer srv.Close()
	td := []struct {
		name    string
		hf      http.HandlerFunc
		pattern string
		v       verifier
	}{
		{"create", srvb.Create, "/CreateURL", creationOK("newgame")},
		{"join", srvb.Join, "/JoinURL", errWith(http.StatusInternalServerError, "no game name was given")},
		{"play", srvb.Play, srvb.PlayURL, errWith(http.StatusInternalServerError, "not created")},
	}
	for _, tc := range td {
		t.Run(fmt.Sprintf("%s endpoint", tc.name), func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			tc.hf(rec, httptest.NewRequest(http.MethodPost, srv.URL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
			res := rec.Result()
			defer res.Body.Close()
			if err := verify(res, tc.v); err != nil {
				t.Fatalf("test failed, err is %v", err)
			}
		})
		srvb.Cleanup(httptest.NewRecorder(), nil)
	}
}

func TestGameRouting(t *testing.T) {
	g := srvb.Game{}
	srv := httptest.NewServer(srvb.Handler())
	defer srv.Close()
	td := []struct {
		name    string
		hf      http.HandlerFunc
		pattern string
		v       verifier
	}{
		{"create", g.Create, srvb.CreateURL, creationOK("newgame")},
		{"join", g.Join, srvb.JoinURL, errWith(http.StatusInternalServerError, "no game name was given")},
		// {"play", srvb.Play, srvb.PlayURL, errWith(http.StatusInternalServerError, "not created")},
	}
	for _, tc := range td {
		t.Run(fmt.Sprintf("%s endpoint", tc.name), func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			tc.hf(rec, httptest.NewRequest(http.MethodPost, srv.URL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
			res := rec.Result()
			defer res.Body.Close()
			if err := verify(res, tc.v); err != nil {
				t.Fatalf("test failed, err is %v", err)
			}
		})
		srvb.Cleanup(httptest.NewRecorder(), nil)
	}
}
