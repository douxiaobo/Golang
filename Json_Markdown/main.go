package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/russross/blackfriday/v2" // 引入Markdown解析库
)

type Article struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	ContentMarkdown string `json:"contentMarkdown"`
	// 可以在这里添加一个字段来存储解析后的HTML，但在实际应用中，通常是在需要时即时解析
}

func main() {
	// 假设你已经从JSON文件中读取了数据并解析到了articles切片中
	// 这里为了简化，我们直接模拟这个过程
	var articles []Article
	// ...（这里应该是读取JSON文件并解析到articles的代码）

	// 假设我们只处理第一篇文章
	if len(articles) > 0 {
		article := articles[0]

		// 解析Markdown为HTML
		htmlOutput := blackfriday.Run([]byte(article.ContentMarkdown))

		// 现在你可以将htmlOutput嵌入到你的HTML模板中，或者直接返回给HTTP客户端
		// 例如，使用http.HandlerFunc来创建一个简单的HTTP服务器
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// 注意：在实际应用中，你应该使用模板来安全地嵌入HTML
			// 这里只是为了演示如何将Markdown解析后的HTML发送到客户端
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, err := w.Write(htmlOutput)
			if err != nil {
				log.Println("Error writing response:", err)
			}
		})

		fmt.Println("Server is listening on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}

// douxiaobo@192 Json_Markdown % go run main.go
// main.go:8:2: no required module provides package github.com/russross/blackfriday/v2: go.mod file not found in current directory or any parent directory; see 'go help modules'
// douxiaobo@192 Json_Markdown % go mod init blackfriday
// go: creating new go.mod: module blackfriday
// go: to add module requirements and sums:
// 	go mod tidy
// douxiaobo@192 Json_Markdown % go get github.com/russross/blackfriday/v2
// go: downloading github.com/russross/blackfriday/v2 v2.1.0
// go: downloading github.com/russross/blackfriday v1.6.0
// go: added github.com/russross/blackfriday/v2 v2.1.0
// douxiaobo@192 Json_Markdown % go run main.go
// douxiaobo@192 Json_Markdown % go mod tidy
// douxiaobo@192 Json_Markdown % go run main.go
// douxiaobo@192 Json_Markdown %
