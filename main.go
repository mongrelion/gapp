package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version string

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer logRequestInfo(r)
		fmt.Fprintf(w, GreetingMessage())
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		defer logRequestInfo(r)
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		defer logRequestInfo(r)
		fmt.Fprintf(w, "ok")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func GetVersion() string {
	return os.Getenv("VERSION")
}

func GreetingMessage() string {
	return fmt.Sprintf("Running version %s", GetVersion())
}

func logRequestInfo(r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL)
}
