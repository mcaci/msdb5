package srvb_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

type ok struct {
	decoder func(io.Reader) (string, error)
	msg     string
}

func (ok) verifyStatusCode(statusCode int) error   { return verifyStatusCode(statusCode, http.StatusOK) }
func (o ok) verifyMessage(resBody io.Reader) error { return verifyMessage(resBody, o.decoder, o.msg) }

func creationOK(msg string) verifier {
	return ok{msg: msg, decoder: func(resBody io.Reader) (string, error) {
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
func joinOK(msg string) verifier {
	return ok{msg: msg, decoder: func(resBody io.Reader) (string, error) {
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

type ko struct {
	statusCode int
	msg        string
}

func (k ko) verifyStatusCode(statusCode int) error { return verifyStatusCode(statusCode, k.statusCode) }
func (k ko) verifyMessage(resBody io.Reader) error {
	return verifyMessage(resBody,
		func(r io.Reader) (string, error) { b, err := ioutil.ReadAll(resBody); return string(b), err },
		k.msg)
}

func errWith(statusCode int, msg string) verifier { return ko{statusCode: statusCode, msg: msg} }

func verifyStatusCode(statusCode1, statusCode2 int) error {
	if statusCode1 != statusCode2 {
		return fmt.Errorf("expected status %d; got %d", statusCode2, statusCode1)
	}
	return nil
}

func verifyMessage(resBody io.Reader, decoder func(io.Reader) (string, error), msg string) error {
	actual, err := decoder(resBody)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}
	expected := msg
	if actual != expected {
		return fmt.Errorf("expecting %q to be in %q", expected, actual)
	}
	return nil
}
