package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
	// 假设URL格式为 /books/someTitle/page/123
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 || parts[1] != "books" || parts[3] != "page" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	title := parts[2]
	pageStr := parts[4]

	// 尝试将pageStr转换为整数，这里仅作示例，实际使用中应处理错误
	page, _ := strconv.Atoi(pageStr)

	fmt.Fprintf(w, "Book Title: %s, Page: %d", title, page)
}

func main() {
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		// 由于/books/后可能跟任意内容，我们在这里重定向到bookHandler，但这种方法比较笨拙
		// 更推荐直接在bookHandler中处理所有逻辑
		bookHandler(w, r)
	})

	// 注意：上面的处理方式并不完美，因为它会匹配所有以/books/开头的路径
	// 更好的做法是直接在HandleFunc中注册具体的路径模式，然后手动解析

	// 更精确的注册方式（但注意，这仍然不是通过{variable}语法实现的）
	// http.HandleFunc("/books/:title/page/:page", bookHandler) // 注意：这不是有效的Go代码

	// 正确的做法是使用自定义的Handler或修改bookHandler来直接处理

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// http://localhost:8080/books/go/page/100
// Book Title: go, Page: 100
