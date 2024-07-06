package main

import (
	"fmt"
	"net/http"
	"strings"
)

var templates = map[string]string{
	"en": `<html><body><h1>Hello, World!</h1></body></html>`,
	"zh": `<html><body><h1>你好，世界！</h1></body></html>`,
	"es": `<html><body><h1>¡Hola, Mundo!</h1></body></html>`,
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 读取Accept-Language头部
	acceptLanguage := r.Header.Get("Accept-Language")

	// 尝试找到最佳匹配的语言
	var lang string
	for _, pref := range strings.Split(acceptLanguage, ",") {
		// 只考虑前两个字符作为语言代码
		langCode := strings.SplitN(pref, ";", 2)[0][:2]
		if _, ok := templates[langCode]; ok {
			lang = langCode
			break
		}
	}

	// 如果没有找到匹配的语言，则默认为英语
	if lang == "" {
		lang = "en"
	}

	// 写入响应内容
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, templates[lang])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain")
// 	fmt.Fprint(w, "Hello World")
// }

func main() {
	// http.HandleFunc("/", handler)

	http.HandleFunc("/", homeHandler)
	fmt.Println("Server is listening on : 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
