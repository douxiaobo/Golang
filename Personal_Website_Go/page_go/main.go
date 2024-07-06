package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// 定义模板文件列表
	templates := []string{"layout.html", "submenu.html", "content.html"}

	// 验证模板文件是否存在
	for _, file := range templates {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Fatalf("Template file %s not found.", file)
		}
	}

	// 加载模板
	tmpl := template.Must(template.ParseFiles(templates...))

	// 设置路由和处理器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 传递数据到模板
		data := struct {
			Title string
		}{
			Title: "我的网页",
		}
		// 直接执行第一个模板（默认是layout.html）
		if err := tmpl.ExecuteTemplate(w, templates[0], data); err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	// 启动服务器
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
