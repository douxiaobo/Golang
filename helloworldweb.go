package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!\n")
}

func main() {
	server := &http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
