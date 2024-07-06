package main

import (
	"net/http"
)

func main() {
	// 设置静态文件的目录
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// 启动服务器
	http.ListenAndServe(":8080", nil)
}
