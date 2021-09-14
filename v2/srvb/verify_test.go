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

type ok struct {
	decoder func(io.Reader) (string, error)
	msg     string
}

func (ok) verifyStatusCode(statusCode int) error { return verifyStatusCode(statusCode, http.StatusOK) }
func (o ok) verifyMessage(resBody io.Reader) error {
	actual, err := o.decoder(resBody)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}
	expected := o.msg
	if actual != expected {
		return fmt.Errorf("expecting %q to be in %q", expected, actual)
	}
	return nil
}

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
	b, err := ioutil.ReadAll(resBody)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}
	actual := string(b)
	expected := k.msg
	if !strings.Contains(actual, expected) {
		return fmt.Errorf("expecting %q to be in %q", expected, actual)
	}
	return nil
}

func errWith(statusCode int, msg string) verifier { return ko{statusCode: statusCode, msg: msg} }

func verifyStatusCode(statusCode1, statusCode2 int) error {
	if statusCode1 != statusCode2 {
		return fmt.Errorf("expected status %d; got %d", statusCode2, statusCode1)
	}
	return nil
}
