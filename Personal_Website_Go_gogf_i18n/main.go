package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gctx"
)

type User struct {
	language string
	name     string
	title    string
}

var languages_ranges = [...]string{"en", "zh-CN", "es"}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var user User
	user.language = getLanguageFromHeader(r.Header.Get("Accept-Language"))
	var (
		ctx  = gctx.New()
		i18n = gi18n.New()
	)

	i18n.SetLanguage(user.language)
	user.title = i18n.Translate(ctx, "title")
	user.name = i18n.Translate(ctx, "name")
	fmt.Fprintf(w, user.title)    //Welcome to my website!
	fmt.Fprintf(w, user.name)     //Douxiaobao
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
	return "en"
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
