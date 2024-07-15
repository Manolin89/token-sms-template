package main

import (
	"log"
	"net/http"
)

func main() {
	s := New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
