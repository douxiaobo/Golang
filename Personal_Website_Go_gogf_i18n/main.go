package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	language string
	name     string
	title    string
}

var languages_ranges = [...]string{"en", "zh", "es"}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var user User
	user.language = getLanguageFromHeader(r.Header.Get("Accept-Language"))
	fmt.Fprintf(w, user.language) //en,zh-CN;q=0.9,zh;q=0.8,es;q=0.7

}

func getLanguageFromHeader(header string) string {
	langs := strings.Split(header, ",")
	for _, lang := range langs {
		trimmedLang := strings.TrimSpace(lang)
		for _, supportedLang := range languages_ranges {
			if strings.HasPrefix(trimmedLang, supportedLang) {
				return supportedLang
			}
		}
	}
	return ""
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
