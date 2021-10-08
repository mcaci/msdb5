package broker_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mcaci/msdb5/v3/srvb/broker"
)

func TestBrokerCreation(t *testing.T) {
	t.Parallel()
	b := broker.NewServer()
	if b == nil {
		t.Error("Could not create broker server")
	}
}

func TestBrokerWithHelloMsg(t *testing.T) {
	t.Parallel()
	b := broker.NewServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	if err != nil {
		t.Error(err)
	}
	rec := httptest.NewRecorder()
	go broker.Listen(b)
	go func() {
		b.Notifier <- []byte("Hello")
	}()
	go time.AfterFunc(350*time.Microsecond, cancel)
	b.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "Hello"
	if !strings.Contains(actual, expected) {
		t.Errorf("expecting %q to be in %q", expected, actual)
	}
}

func TestBrokerWithTimeMsg(t *testing.T) {
	t.Parallel()
	b := broker.NewServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	if err != nil {
		t.Error(err)
	}
	rec := httptest.NewRecorder()
	var count int
	go broker.Listen(b)
	go func() {
		ticker := time.NewTicker(100 * time.Microsecond)
		go time.AfterFunc(350*time.Microsecond, func() { ticker.Stop(); cancel() })
		for tick := range ticker.C {
			count++
			eventString := fmt.Sprintf("the time is %v", tick.Local())
			t.Log(eventString)
			b.Notifier <- []byte(eventString)
		}
	}()
	b.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	t.Log(actual)
	if count > 3 {
		t.Error("expecting the count of ticks to be 3 or less")
	}
}
