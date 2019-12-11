package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexRouter)
	http.Handle("/", r)
}

func IndexRouter(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello World")
}
