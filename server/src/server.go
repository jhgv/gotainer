package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! I'm a HTTP server!")
	})

	var port = ":8888"
	if value, ok := os.LookupEnv("PORT"); ok {
		port = fmt.Sprintf(":%s", value)
	}
	log.Fatal(http.ListenAndServe(port, nil))
}
