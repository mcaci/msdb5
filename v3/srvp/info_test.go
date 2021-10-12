package srvp_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mcaci/msdb5/v3/srvp"
)

func TestBasicInfo(t *testing.T) {
	in := fmt.Sprintf(`{"inturn":%t}`, true)
	req := httptest.NewRequest(http.MethodPost, "/info", strings.NewReader(in))
	rec := httptest.NewRecorder()
	srvp.Info(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("Expecting 200 but got", res.StatusCode)
	}
	var out struct {
		InTurn bool `json:"inturn"`
	}
	err := json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		t.Fatal(err)
	}
	if !out.InTurn {
		t.Error("Expecting your turn to be true but was false")
	}
}

func TestInputError(t *testing.T) {
	in := fmt.Sprintf(`{"inturn":"%t"}`, true)
	req := httptest.NewRequest(http.MethodPost, "/info", strings.NewReader(in))
	rec := httptest.NewRecorder()
	srvp.Info(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expecting %d but got %d", http.StatusBadRequest, res.StatusCode)
	}
}
