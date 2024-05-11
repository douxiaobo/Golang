package main

import (
	"fmt"
	"net/http"
)

// IndexHandler 处理根路径的请求，返回带有"Hello, Go!"的HTML页面
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应的HTTP头部，指定返回的内容类型是HTML
	w.Header().Set("Content-Type", "text/html")

	// 写入响应体，即HTML内容
	html := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Hello, Go!</title>
</head>
<body>
    <h1>Hello, Go!</h1>
</body>
</html>
`
	// 将HTML内容写入HTTP响应
	w.Write([]byte(html))
}

func main() {
	// 设置路由，将根路径("/")映射到IndexHandler函数
	http.HandleFunc("/", IndexHandler)

	// 启动HTTP服务器，监听8080端口
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// 如果启动服务器失败，打印错误信息
		fmt.Println(err)
	}
}

// 这段代码定义了一个IndexHandler函数，该函数在接收到HTTP请求时被调用，并且返回一个HTML页面。
// 然后，main函数中设置了根路径的路由，并启动了一个监听8080端口的HTTP服务器。
// 当你在浏览器中访问http://localhost:8080时，你将看到显示"Hello, Go!"的网页。

// 请确保你的Go环境已经设置好，并且你已经安装了必要的依赖。
// 运行这段代码，然后在浏览器中访问指定的URL，你应该能看到"Hello, Go!"的显示。
