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

func TestPlayWithNoBodyError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, srvb.JoinURL, nil)
	if err != nil {
		t.Error(err)
	}
	g.Play(rec, req)
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

func Test1PPlay(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame"))))
	rec := httptest.NewRecorder()
	g.Play(rec, httptest.NewRequest(http.MethodPost, srvb.PlayURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", "newgame", 1))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("got", res.StatusCode)
	}
	var rs struct {
		Pl  string `json:"player"`
		Brd string `json:"board"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		t.Error(err)
	}
	s := fmt.Sprint(rs)
	if !(strings.Contains(s, "Name") && strings.Contains(s, "Cards")) {
		t.Error("expecting this gameplay result: ", s)
	}
	g.Cleanup()
}

func Test2PPlay(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame"))))
	g.Play(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.PlayURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", "newgame", 1))))
	rec := httptest.NewRecorder()
	g.Play(rec, httptest.NewRequest(http.MethodPost, srvb.PlayURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "michi", "newgame", 0))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("got", res.StatusCode)
	}
	var rs struct {
		Pl  string `json:"player"`
		Brd string `json:"board"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		t.Error(err)
	}
	s := fmt.Sprint(rs)
	if !(strings.Contains(s, "Name") && strings.Contains(s, "Cards")) {
		t.Error("expecting this gameplay result: ", s)
	}
	g.Cleanup()
}

func Test1PPlayTwiceError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	g.Create(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.CreateURL, strings.NewReader(fmt.Sprintf(`{"name":"%s"}`, "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "mary", "newgame"))))
	g.Join(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.JoinURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s"}`, "michi", "newgame"))))
	g.Play(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, srvb.PlayURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", "newgame", 1))))
	rec := httptest.NewRecorder()
	g.Play(rec, httptest.NewRequest(http.MethodPost, srvb.PlayURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", "newgame", 0))))
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("got", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "not expected to play"
	if !strings.Contains(string(b), expected) {
		t.Error("expecting to have", expected, "in", string(b))
	}
	g.Cleanup()
}

func TestPlayOnNotExistingGameError(t *testing.T) {
	t.Parallel()
	g := srvb.Game{}
	rec := httptest.NewRecorder()
	g.Play(rec, httptest.NewRequest(http.MethodPost, srvb.PlayURL, strings.NewReader(fmt.Sprintf(`{"name":"%s","game":"%s","id":%d}`, "mary", "newgame", 1))))
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
