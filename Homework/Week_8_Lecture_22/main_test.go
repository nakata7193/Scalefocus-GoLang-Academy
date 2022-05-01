package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockServer () *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
	}))
}
