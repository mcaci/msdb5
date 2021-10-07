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

	"github.com/mcaci/msdb5/v3/srvb"
)

func TestRouting(t *testing.T) {
	t.Parallel()
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
		{"play", g.Play, srvb.PlayURL, errWith(http.StatusInternalServerError, "not created")},
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
		g.Cleanup()
	}
}

type verifier interface {
	verifyStatusCode(statusCode int) error
	verifyMessage(resBody io.Reader) error
}

func verify(res *http.Response, v verifier) error {
	defer res.Body.Close()
	if err := v.verifyStatusCode(res.StatusCode); err != nil {
		b, errb := ioutil.ReadAll(res.Body)
		if errb != nil {
			return fmt.Errorf("could not read response body: %v", errb)
		}
		return fmt.Errorf("got different error code: %v. Cause is: %v", err, string(b))
	}
	if err := v.verifyMessage(res.Body); err != nil {
		return err
	}
	return nil
}

func creationOK(msg ...string) *expectedData {
	return &expectedData{statusCode: http.StatusOK, msg: msg, decoder: func(resBody io.Reader) (string, error) {
		var rs struct {
			Name string `json:"name"`
		}
		err := json.NewDecoder(resBody).Decode(&rs)
		if err != nil {
			return "", err
		}
		return rs.Name, nil
	}}
}

func errWith(statusCode int, msg ...string) *expectedData {
	return &expectedData{statusCode: statusCode, msg: msg, decoder: func(resBody io.Reader) (string, error) {
		b, err := ioutil.ReadAll(resBody)
		if err != nil {
			return "", err
		}
		return string(b), err
	}}
}

type expectedData struct {
	statusCode int
	decoder    func(io.Reader) (string, error)
	msg        []string
}

func (ed *expectedData) verifyStatusCode(statusCode int) error {
	if ed.statusCode != statusCode {
		return fmt.Errorf("expected status %d; got %d", ed.statusCode, statusCode)
	}
	return nil
}
func (ed *expectedData) verifyMessage(resBody io.Reader) error {
	actual, err := ed.decoder(resBody)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}
	expecteds := ed.msg
	for _, expected := range expecteds {
		if !strings.Contains(actual, expected) {
			return fmt.Errorf("expecting %q to be in %q", expected, actual)
		}
	}
	return nil
}
