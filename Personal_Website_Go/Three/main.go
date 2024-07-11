package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var lang = []string{"en", "zh", "es"}
	// 读取Accept-Language头部
	// acceptLanguage := r.Header.Get("Accept-Language")

	// var langCode string

	// langCode = strings.Split(acceptLanguage, ",")[0][:2]
	langCode := strings.Split(r.Header.Get("Accept-Language"), ",")[0][:2]
	// fmt.Println("Trying language", langCode)

	found := false
	for _, l := range lang {
		if l == langCode {
			found = true
			break
		}
	}
	if !found {
		langCode = "en"
	}

	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalf("Failed to get working directory: %v", err)
	// }
	// fmt.Println("Current Working Directory:", dir)

	templates := []string{"index.html", "submenu.html", "content.html"}

	// url := filepath.Join(dir, "public/tmpl/", langCode)
	// 定义模板文件列表
	// templates := []string{url + "/index.html", url + "/submenu.html", url + "/content.html"}
	// fmt.Println(templates)
	// templates := []string{
	// 	filepath.Join(dir, "public/tmpl", langCode, "index.html"),
	// 	filepath.Join(dir, "public/tmpl", langCode, "submenu.html"),
	// 	filepath.Join(dir, "public/tmpl", langCode, "content.html"),
	// }

	// 验证模板文件是否存在
	for _, file := range templates {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Fatalf("Template file %s not found.", file)
		}
	}

	// fmt.Println(templates)

	// 加载模板
	tmpl := template.Must(template.ParseFiles(templates...))
	fmt.Println("Loaded templates:", tmpl.Templates())

	data := struct {
		Title string
	}{
		Title: "我的主页",
	}

	if err := tmpl.ExecuteTemplate(w, templates[0], data); err != nil {
		log.Printf("Error executing template: %v", err)
		// http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", homeHandler)

	//启动服务器
	// fmt.Println("Server is listening on :8080")
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Could not start server: %v", err)
	}
}
