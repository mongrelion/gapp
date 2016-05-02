package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version string
var bind string

func init() {
	version = os.Getenv("VERSION")
	if len(version) == 0 {
		version = "unknown version"
	}

	bind = os.Getenv("BIND")
	if len(bind) == 0 {
		bind = ":8080"
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, version)
	})
	fmt.Printf("Binding to %s\n", bind)
	fmt.Printf("Running version %s\n", version)
	log.Fatal(http.ListenAndServe(bind, nil))
}
