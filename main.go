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
		fmt.Fprintf(w, GreetingMessage())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetVersion() string {
	return os.Getenv("VERSION")
}

func GreetingMessage() string {
	return fmt.Sprintf("Running version %s", GetVersion())
}
