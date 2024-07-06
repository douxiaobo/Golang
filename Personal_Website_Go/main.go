package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var templates map[string]*template.Template

func initTemplates() {
	templates = make(map[string]*template.Template)
	// templates["en"], _ = template.ParseFiles("public/tmpl/en/index.html")
	// templates["zh"], _ = template.ParseFiles("public/tmpl/zh/index.html")
	// templates["es"], _ = template.ParseFiles("public/tmpl/es/index.html")

	var baseTemplate *template.Template
	var err error
	if baseTemplate, err = template.ParseFiles("public/tmpl/base.html"); err != nil {
		log.Fatalf("Error parsing base template: %v", err)
	}

	for lang, path := range map[string]string{
		"en": "public/tmpl/en/index.html",
		"zh": "public/tmpl/zh/index.html",
		"es": "public/tmpl/es/index.html",
	} {
		templates[lang] = template.Must(baseTemplate.Clone())
		templates[lang], err = templates[lang].ParseFiles(path)
		if err != nil {
			log.Fatalf("Error parsing language template for %s: %v", lang, err)
		}
		templates[lang], err = templates[lang].ParseGlob("public/tmpl/*.html")
		if err != nil {
			log.Fatalf("Error parsing glob templates for %s: %v", lang, err)
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 读取Accept-Language头部
	acceptLanguage := r.Header.Get("Accept-Language")

	// 尝试找到最佳匹配的语言
	for _, pref := range strings.Split(acceptLanguage, ",") {
		// 只考虑前两个字符作为语言代码
		langCode := strings.SplitN(pref, ";", 2)[0][:2]
		if tmpl, ok := templates[langCode]; ok {
			// 执行模板渲染
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	// 如果没有找到匹配的语言，则默认为英语
	if tmpl, ok := templates["en"]; ok {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// 如果连英语模板都没有，就返回一个错误
		http.Error(w, "No template found", http.StatusInternalServerError)
	}
}

func main() {
	initTemplates()
	http.HandleFunc("/", homeHandler)

	//启动服务器
	// fmt.Println("Server is listening on :8080")
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Could not start server: %v", err)
	}
}
