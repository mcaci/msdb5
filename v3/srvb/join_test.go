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

func Test1PJoin(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("got", res.StatusCode)
	}
	var rs struct {
		Number string `json:"number"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		t.Error(err)
	}
	if rs.Number != "1" {
		t.Error("expecting a game with name", rs.Number)
	}
	g.Cleanup()
}

func Test2PJoin(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame"))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("got", res.StatusCode)
	}
	var rs struct {
		Number string `json:"number"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		t.Error(err)
	}
	if rs.Number != "2" {
		t.Error("expecting a game with name", rs.Number)
	}
	g.Cleanup()
}

func Test3PJoinError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "thirdplayer", "newgame"))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "max players reached"
	if !strings.Contains(string(b), expected) {
		t.Error("expecting to have", expected, "in", string(b))
	}
	g.Cleanup()
}

func TestWrongBodyJoinError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(`'{"name":"na"}`)))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "invalid character"
	if !strings.Contains(string(b), expected) {
		t.Error("expecting to have", expected, "in", string(b))
	}
	g.Cleanup()
}

// {"Join with no body gives error", []*operation{
// 	join(nil),
// }, errWith(http.StatusBadRequest, "empty request")},
// {"Join with no create gives error", []*operation{
// 	join(defaultGame("mary")),
// }, errWith(http.StatusInternalServerError, "not created")},
// {"Join on wrong game", []*operation{
// 	create(withName(newgame)),
// 	join(strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame"))),
// }, errWith(http.StatusInternalServerError, "different name")},
// {"Join with no player name gives error", []*operation{
// 	create(withName(newgame)),
// 	join(strings.NewReader(fmt.Sprintf(`{"game":"%s"}`, newgame))),
// }, errWith(http.StatusInternalServerError, "no player name was given")},
// {"Join with no game name gives error", []*operation{
// 	create(withName(newgame)),
// 	join(strings.NewReader(`{"name":"player"}`)),
// }, errWith(http.StatusInternalServerError, "no game name was given")},
