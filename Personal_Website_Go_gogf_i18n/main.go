package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gctx"
)

type User struct {
	Language    string
	Name        string
	Title       string
	ContentName string
	FooterLinks []FooterLink
}

type FooterLink struct {
	Lang                  string
	Footer_language_short string
	Footer_language_long  string
}

var languages_ranges = [...]string{"en", "zh", "es"}

var containtlist = [...]string{"home", "about", "work", "travel", "music", "programming", "school", "sport"}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var user User
	// 从URL路径中获取语言后缀
	path := strings.TrimPrefix(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")

	// // 如果路径为空或仅为 "/"
	// if len(pathParts) == 0 || (len(pathParts) == 1 && pathParts[0] == "") {
	// 	// 设置默认语言和内容名称
	// 	user.Language = getLanguageFromHeader(r.Header.Get("Accept-Language"))
	// 	user.ContentName = "home"
	// }

	// // 如果路径中有足够的部分，那么第一部分是语言，第二部分是内容名称
	// if len(pathParts) >= 2 {
	// 	user.Language = pathParts[0]
	// 	user.ContentName = pathParts[1]

	// 	if !(isValidLang(pathParts[0]) || isValidContentName(pathParts[1])) {
	// 		http.Error(w, "Invaid URL", http.StatusBadGateway)
	// 		return
	// 	}
	// } else if len(pathParts) >= 2 && pathParts[1] == "" {
	// 	user.Language = pathParts[0]
	// 	user.ContentName = "home"
	// } else if len(pathParts) == 1 {
	// 	user.Language = pathParts[0]
	// 	// 如果没有指定内容名称，默认使用home
	// 	if !isValidLang(user.Language) {
	// 		http.Error(w, "Invalid URL", http.StatusBadRequest)
	// 		return
	// 	}
	// 	// user.ContentName = "home"
	// } else {
	// 	user.Language = getLanguageFromHeader(r.Header.Get("Accept-Language"))
	// 	// user.ContentName = "home" // 这里替换为你的默认内容名称
	// }

	// 如果路径为空或仅为 "/"
	if len(pathParts) == 0 || (len(pathParts) == 1 && pathParts[0] == "") {
		// 设置默认语言和内容名称
		user.Language = getLanguageFromHeader(r.Header.Get("Accept-Language"))
		user.ContentName = "home"
	} else if len(pathParts) == 1 || len(pathParts) == 2 && pathParts[1] == "" {
		user.Language = pathParts[0]
		user.ContentName = "home"
	} else if len(pathParts) == 2 {
		user.Language = pathParts[0]
		user.ContentName = pathParts[1]
	} else {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// 检查语言是否在支持的范围内
	if !isValidLang(user.Language) {
		user.Language = "en"
		http.Redirect(w, r, fmt.Sprintf("/%s/%s", user.Language, user.ContentName), http.StatusMovedPermanently)
		return
	}

	if !isValidContentName(user.ContentName) {
		user.ContentName = "home"
		http.Redirect(w, r, fmt.Sprintf("/%s/%s", user.Language, user.ContentName), http.StatusMovedPermanently)
		return
	}

	var (
		ctx  = gctx.New()
		i18n = gi18n.New()
	)

	i18n.SetLanguage(user.Language)
	user.Title = i18n.Translate(ctx, "title")
	user.Name = i18n.Translate(ctx, "name")

	// t, err := template.ParseFiles("templates/index.tmpl", "templates/nav.tmpl", "templates/header.tmpl", "templates/footer.tmpl")
	// t, err := template.ParseFiles("templates/index.tmpl")
	t, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		fmt.Println("Template parsing error:", err)
		return
	}
	if err := t.ExecuteTemplate(w, "index.tmpl", user); err != nil {
		log.Println("Template execution error: ", err)
		return
	}

	// fmt.Fprintf(w, user.title)    //Welcome to my website!
	// fmt.Fprintf(w, user.name)     //Douxiaobao
	// fmt.Fprintf(w, user.language) //en,zh-CN;q=0.9,zh;q=0.8,es;q=0.7

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

func isValidLang(lang string) bool {
	for _, l := range languages_ranges {
		if lang == l {
			return true
		}
	}
	return false
}

func isValidContentName(contentName string) bool {
	for _, c := range containtlist {
		if contentName == c {
			return true
		}
	}
	return false
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
