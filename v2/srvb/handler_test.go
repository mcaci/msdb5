package srvb_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mcaci/msdb5/v2/srvb"
)

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(srvb.Handler())
	defer srv.Close()

	res, err := http.Post(fmt.Sprintf("%s/create", srv.URL), "application/json", strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame")))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", res.StatusCode)
	}
	actual, err := func(res *http.Response) (string, error) {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}(res)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	expected := "{\"name\":\"newgame\"}\n"
	if actual != expected {
		t.Fatalf("expecting %q got %q", expected, actual)
	}
}
