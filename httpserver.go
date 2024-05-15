package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var data = map[string]string{"testweb": "hello world！"} // 网页界面显示信息

// 定义一个处理HTTP GET请求的处理器
func handleGet(w http.ResponseWriter, r *http.Request) {
	// 发送HTTP 200响应
	w.WriteHeader(http.StatusOK)
	// 设置响应头Content-Type为application/json
	w.Header().Set("Content-Type", "application/json")
	// 序列化data为JSON并写入响应体
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	host := "localhost"
	port := "8080"
	// 设置监听地址和端口
	addr := fmt.Sprintf("%s:%s", host, port)

	// 注册HTTP GET请求的处理函数
	http.HandleFunc("/", handleGet)

	// 启动HTTP服务器
	fmt.Printf("Starting server, listen at: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// 在这个Go程序中，我们定义了一个handleGet函数来处理HTTP GET请求。当请求到达时，这个函数会被调用，并生成一个JSON响应。
// 我们使用json.Marshal函数来序列化data映射为JSON格式的字节切片，并将其写入响应体。

// 在main函数中，我们设置了监听地址和端口，并通过http.HandleFunc函数将/路径映射到handleGet处理器函数。然后，我们使用http.ListenAndServe函数启动HTTP服务器。
// 这个函数会阻塞，直到服务器停止。如果服务器在监听时发生错误，log.Fatal会记录错误并退出程序。
