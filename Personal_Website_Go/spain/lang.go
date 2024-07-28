package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	lang := "zh"
	path := r.URL.Path
	if strings.HasPrefix(path, "/en/") {
		lang = "en"
	}

	var title, content string
	if lang == "zh" {
		title = "窦小波的个人网站"
		content = "嗨，欢迎你来访问我的个人主页，谢谢！"
	} else if lang == "en" {
		title = "Dou Xiaobo's Personal Website"
		content = "Hi, welcome to my personal homepage, thank you!"
	}

	fmt.Fprintf(w, "Language: %s\n", lang)
	fmt.Fprintf(w, "Title: %s\n", title)
	fmt.Fprintf(w, "Content: %s\n", content)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/en/", homeHandler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server at port 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
