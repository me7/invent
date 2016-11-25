package main

import (
	"fmt"
	"log"
	"net/http"
)

type mainHandler struct{}

func main() {
	var mainHandler mainHandler
	m := http.NewServeMux()

	m.Handle("/", mainHandler)
	log.Fatal(http.ListenAndServe(":7749", m))
}

func (h mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you're come from %s", r.RemoteAddr)
}
