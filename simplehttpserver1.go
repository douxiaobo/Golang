package main

import (
	"fmt"
	"net/http"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/helloworld", SayHelloWorld)
	http.ListenAndServe(":8080", nil)
}
