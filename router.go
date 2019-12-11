package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter initializes a new router, or map, for any request incoming to the API.
// NewRouter works as a set of instructions, matching an endpoint that Slab will listen for to
// a function that should be ran when that endpoint is requested.
func NewRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexRouter)
	http.Handle("/", r)
}

// IndexRouter is the root level endpoint that's returned when a user requests the "/" endpoint.
// This route will likely be utilized in the future to show the current server status.
func IndexRouter(w http.ResponseWriter, r *http.Request) {
	// Using JSON, set a very basic health check
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO: add in additional health metrics, such as uptime, database connections, etc.
	io.WriteString(w, `{"alive": true}`)
}
