package srvb_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mcaci/msdb5/v2/srvb"
)

func TestCreation(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/create", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	srvb.Create(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	expected := "OK"
	if string(b) != expected {
		t.Fatalf("expecting %v got %v", expected, string(b))
	}
}

// TODO
func TestCreationWithName(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/create?name=new", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	srvb.Create(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	expected := "OK"
	if string(b) != expected {
		t.Fatalf("expecting %v got %v", expected, string(b))
	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(srvb.Handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/create", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	expected := "OK"
	if string(b) != expected {
		t.Fatalf("expecting %v got %v", expected, string(b))
	}

}
