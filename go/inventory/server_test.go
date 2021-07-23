package inventory

import (
	"net/http"
	"testing"
	"time"
)

func TestHealth(t *testing.T) {
	s := NewServer()

	http_s := &http.Server{
		Addr:           ":8080",
		Handler:        s.engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	defer http_s.Close()

	// Server requests in go routine
	go http_s.ListenAndServe()

	time.Sleep(2)

	res, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error("invalid returned status")
	}
}
