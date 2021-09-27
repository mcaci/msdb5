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
		pattern string
		v       verifier
	}{
		{"create", srvb.CreateURL, creationOK("newgame")},
		{"join", srvb.JoinURL, errWith(http.StatusInternalServerError, "no game name was given")},
		{"play", srvb.PlayURL, errWith(http.StatusInternalServerError, "not created")},
	}
	for _, tc := range td {
		t.Run(fmt.Sprintf("test %s endpoint", tc.name), func(t *testing.T) {
			res, err := http.Post(fmt.Sprintf("%s%s", srv.URL, tc.pattern), "application/json", strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")))
			if err != nil {
				t.Fatalf("could not send POST request: %v", err)
			}
			defer res.Body.Close()
			if err := verify(res, tc.v); err != nil {
				t.Fatalf("test failed, err is %v", err)
			}
			srvb.Cleanup(httptest.NewRecorder(), nil)
		})
	}
}
