package srvb_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const host = "localhost:8080"

func testResWith(expectedStatusCode int, decoder func(res *http.Response) (string, error)) func(*http.Response, string) error {
	cmp := strings.Contains
	if expectedStatusCode == http.StatusOK {
		cmp = func(a, b string) bool { return a == b }
	}
	return func(res *http.Response, expected string) error {
		if res.StatusCode != expectedStatusCode {
			return fmt.Errorf("expected status %d; got %d", expectedStatusCode, res.StatusCode)
		}
		actual, err := decoder(res)
		if err != nil {
			return fmt.Errorf("could not read response: %v", err)
		}
		if !cmp(actual, expected) {
			return fmt.Errorf("expecting %q to be in %q", expected, actual)
		}
		return nil
	}
}

func koDec(res *http.Response) (string, error) {
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("could not read error response: %v", err)
	}
	return string(b), nil
}
