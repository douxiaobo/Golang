package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// 假设你有实际的模板文件：templates/index_en.html, templates/index_zh.html, templates/index_es.html

var templates map[string]*template.Template

func initTemplates() {
	templates = make(map[string]*template.Template)
	// 加载模板文件到templates map中
	templates["en"], _ = template.ParseFiles("templates/index_en.html")
	templates["zh"], _ = template.ParseFiles("templates/index_zh.html")
	templates["es"], _ = template.ParseFiles("templates/index_es.html")
	// 注意：在实际应用中，你应该处理这些ParseFiles调用可能返回的错误
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 读取Accept-Language头部
	acceptLanguage := r.Header.Get("Accept-Language")

	// 尝试找到最佳匹配的语言
	for _, pref := range strings.Split(acceptLanguage, ",") {
		// 只考虑前两个字符作为语言代码
		langCode := strings.SplitN(pref, ";", 2)[0][:2]
		if tmpl, ok := templates[langCode]; ok {
			// 执行模板渲染
			tmpl.Execute(w, nil) // 假设模板不需要任何数据
			return
		}
	}

	// 如果没有找到匹配的语言，则默认为英语
	if tmpl, ok := templates["en"]; ok {
		tmpl.Execute(w, nil) // 使用英语模板
	} else {
		// 如果连英语模板都没有，就返回一个错误
		http.Error(w, "No template found", http.StatusInternalServerError)
	}
}

func main() {
	initTemplates() // 初始化模板
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
