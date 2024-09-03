package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	language string
	name     string
	title    string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var user User
	user.language = getLanguageFromHeader(r)
	fmt.Fprintf(w, user.language) //en,zh-CN;q=0.9,zh;q=0.8,es;q=0.7

}

func getLanguageFromHeader(r *http.Request) string {
	return r.Header.Get("Accept-Language")
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in main: %v", r)
		}
	}()
	log.Println("Starting HTTP server...")
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
