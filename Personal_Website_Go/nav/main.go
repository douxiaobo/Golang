package main

import (
	"html/template"
	"net/http"
)

func main() {
	t := template.New("base").Funcs(template.FuncMap{
		"httphead": func() string { return `` }, // 确保函数名是小写的
	})

	// 确保 templates 目录下的模板文件被解析
	templates, err := t.ParseGlob("templates/*")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "about.html", nil) // 使用 ExecuteTemplate 执行 about.html
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
