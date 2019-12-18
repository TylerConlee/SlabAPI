package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var gettests = []struct {
	name     string
	route    string
	handler  http.HandlerFunc
	expected string
}{
	{"GET - IndexHandler", "/", IndexRouter, `{"alive": true}`},
	{"GET - ConfigHandler", "/config", GetConfigHandler, `{"alive": false}`},
}

func TestGETHandlers(t *testing.T) {
	for _, tt := range gettests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
			// pass 'nil' as the third parameter.
			req, err := http.NewRequest("GET", tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tt.handler)

			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// Check the response body is what we expect.

			if rr.Body.String() != tt.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expected)
			}
		})
	}
}
