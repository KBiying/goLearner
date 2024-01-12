package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
func ping(url string) chan bool {
	ch := make((chan bool))
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

var timeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return CongrableRacer(a, b, timeout)
}
func CongrableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): //send a signal into channel after 10s
		return "", fmt.Errorf("timed out wating for %s and %s", a, b)
	}
}

// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b
// }

func TestRacer(t *testing.T) {
	t.Run("select quiker one", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("fid not eacpetc an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("return while out of time 10s ", func(t *testing.T) {
		serverA := makeDelayServer(11 * time.Second)
		serverB := makeDelayServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := CongrableRacer(serverA.URL, serverB.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})
}
