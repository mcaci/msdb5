package srvb_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mcaci/msdb5/v3/srvb"
)

func TestCreateObj(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	rec := httptest.NewRecorder()
	// curl -XPOST  -H "Content-Type: application/json" localhost:8080/create -d '{"name":"newgame"}'
	g.Create(rec, httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	if g.Game == nil {
		t.Error("error")
	}
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("got", res.StatusCode)
	}
	var rs struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		t.Error(err)
	}
	if rs.Name == "" {
		t.Error("expecting a game with name", rs.Name)
	}
	g.Cleanup()
}

func TestCreateWithNoBodyGivesErr(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, srvb.CreateURL, nil)
	if err != nil {
		t.Error(err)
	}
	g.Create(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) == "" {
		t.Error("expecting to have:", string(b))
	}
	g.Cleanup()
}

func TestCreateWithErrBodyGivesErr(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(`'{"name":"na"}`))
	g.Create(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) == "" {
		t.Error("expecting to have:", string(b))
	}
	g.Cleanup()
}

func TestCannotCreateTwoGames(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	rec := httptest.NewRecorder()
	g.Create(rec, httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "othergame"))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(b), "one game already created, cannot create more") {
		t.Error("expecting to have 'one game already created, cannot create more' in", string(b))
	}
	g.Cleanup()
}
