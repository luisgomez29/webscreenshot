package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScreenshotHandler(t *testing.T) {
	// Table driven test
	tt := []struct {
		url        string
		statusCode int
	}{
		{"https://www.google.com", http.StatusOK},
		{"https://www.facebook.com", http.StatusOK},
		{"https://www.youtube.com", http.StatusOK},
		{"https://www.pinterest.com", http.StatusOK},
		{"https://www.tripadvisor.com", http.StatusOK},
		{"https://www.yahoo.com", http.StatusOK},
		{"https://www.twitter.com", http.StatusOK},
		{"https://www.instagram.com", http.StatusOK},
		{"https://www.github.com", http.StatusOK},
		{"https://www.gitlab.com", http.StatusOK},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/screenshot?url=%s", tc.url)
		req := httptest.NewRequest("GET", path, nil)
		rr := httptest.NewRecorder()

		// need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/screenshot", ScreenshotHandler)
		router.ServeHTTP(rr, req)

		// check response status
		if rr.Code != tc.statusCode {
			t.Errorf("handler should have failed on url %s: got %v want %v", tc.url, rr.Code, tc.statusCode)
		}
	}
}
