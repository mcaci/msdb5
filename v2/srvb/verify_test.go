package srvb_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type crOK string

func (c crOK) verify(res *http.Response) error {
	return okVerify(res, creationDec, string(c))
}

type joinOK string

func (j joinOK) verify(res *http.Response) error {
	return okVerify(res, joinDec, string(j))
}

func okVerify(res *http.Response, decoder func(res *http.Response) (string, error), msg string) error {
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status %d; got %d", http.StatusOK, res.StatusCode)
	}
	actual, err := decoder(res)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}
	expected := msg
	if actual != expected {
		return fmt.Errorf("expecting %q to be in %q", expected, actual)
	}
	return nil
}

type badReqErr string

func (b badReqErr) verify(res *http.Response) error {
	return koVerify(res, http.StatusBadRequest, string(b))
}

type intSrvErr string

func (i intSrvErr) verify(res *http.Response) error {
	return koVerify(res, http.StatusInternalServerError, string(i))
}

func koVerify(res *http.Response, statusCode int, msg string) error {
	if res.StatusCode != statusCode {
		return fmt.Errorf("expected status %d; got %d", statusCode, res.StatusCode)
	}
	actual, err := koDec(res)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}
	expected := msg
	if !strings.Contains(actual, expected) {
		return fmt.Errorf("expecting %q to be in %q", expected, actual)
	}
	return nil
}

func creationDec(res *http.Response) (string, error) {
	var rs struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		return "", err
	}
	return rs.Name, nil
}

func joinDec(res *http.Response) (string, error) {
	var rs struct {
		Number string `json:"number"`
	}
	err := json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		return "", err
	}
	return rs.Number, nil
}

func koDec(res *http.Response) (string, error) {
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("could not read error response: %v", err)
	}
	return string(b), nil
}
