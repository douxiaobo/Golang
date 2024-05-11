package main

import (
	"fmt"
	"net/http"
)

// 生成响应文本
func makeResponse() string {
	response := "HTTP/1.1 200 OK\r\n"
	response += "Server: My Go Server\r\n"
	response += "Content-Type: text/html\r\n"
	response += "\r\n<!DOCTYPE html>\r\n\r\n"
	response += "<html>"
	response += "<head><meta charset=\"utf-8\"></head>"
	response += "<body><h1>Hello, Go!</h1></body>"
	response += "</html>\r\n"
	return response
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	// 输出来自浏览器的请求信息
	fmt.Printf("\"%s\"\n", req.RequestURI)

	// 发送响应文档
	response := makeResponse()
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", handleRequest) // 设置处理函数和路由

	// 启动服务器
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

// 这段 Go 代码使用了 net/http 包来创建 HTTP 服务器，并通过 http.HandleFunc 设置了路由和处理函数。
// 它没有直接使用 net 包来处理底层的 TCP 连接和读取请求，因为 net/http 已经提供了更高层次的封装。
// 如果你需要更精细的控制，比如直接读取请求头，你可能需要使用 net 包，但这通常不是处理 HTTP 请求的推荐方式。

// 请注意，Go 标准库中的 net/http 已经足够强大，可以处理大多数的 HTTP 服务器需求，因此通常不需要直接操作 TCP 套接字。
// 此外，Go 语言的 http.ResponseWriter 接口提供了一种方便的方式来发送响应给客户端。
