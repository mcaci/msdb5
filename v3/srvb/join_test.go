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
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
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
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame"))))
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
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "thirdplayer", "newgame"))))
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

func TestJoinWithNoBodyGivesErr(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, srvb.JoinURL, nil)
	if err != nil {
		t.Error(err)
	}
	g.Join(rec, req)
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

func TestWrongBodyJoinError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(`'{"name":"na"}`)))
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

func TestJoinOnWrongGameError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "othergame"))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "different name"
	if !strings.Contains(string(b), expected) {
		t.Error("expecting to have", expected, "in", string(b))
	}
	g.Cleanup()
}

func TestJoinWithNoGameCreatedGivesError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "not created"
	if !strings.Contains(string(b), expected) {
		t.Error("expecting to have", expected, "in", string(b))
	}
	g.Cleanup()
}

func TestJoinWithNoGameNameGivesError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "mary"))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "no game name was given"
	if !strings.Contains(string(b), expected) {
		t.Error("expecting to have", expected, "in", string(b))
	}
	g.Cleanup()
}

func TestJoinWithNoPlayerNameGivesError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	rec := httptest.NewRecorder()
	g.Join(rec, httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"game":"%s"}`, "newgame"))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "no player name was given"
	if !strings.Contains(string(b), expected) {
		t.Error("expecting to have", expected, "in", string(b))
	}
	g.Cleanup()
}
