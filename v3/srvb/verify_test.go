package srvb_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type verifier interface {
	verifyStatusCode(statusCode int) error
	verifyMessage(resBody io.Reader) error
}

func verify(res *http.Response, v verifier) error {
	if err := v.verifyStatusCode(res.StatusCode); err != nil {
		return err
	}
	defer res.Body.Close()
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
func joinOK(msg ...string) *expectedData {
	return &expectedData{statusCode: http.StatusOK, msg: msg, decoder: func(resBody io.Reader) (string, error) {
		var rs struct {
			Number string `json:"number"`
		}
		err := json.NewDecoder(resBody).Decode(&rs)
		if err != nil {
			return "", err
		}
		return rs.Number, nil
	}}
}
func playOK(msg ...string) *expectedData {
	return &expectedData{statusCode: http.StatusOK, msg: msg, decoder: func(resBody io.Reader) (string, error) {
		var rs struct {
			Pl  string `json:"player"`
			Brd string `json:"board"`
		}
		err := json.NewDecoder(resBody).Decode(&rs)
		if err != nil {
			return "", err
		}
		return fmt.Sprint(rs), nil
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
