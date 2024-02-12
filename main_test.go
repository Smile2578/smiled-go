// main_test.go

package main

import (
	"net/http"
	"testing"
	"time"
)

func TestServerStart(t *testing.T) {
	go func() {
		main()
	}()

	time.Sleep(time.Second) // Wait for the server to start

	_, err := http.Get("http://localhost:8080") // Replace with your server's address
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
}
