package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 创建一个HTTP请求
	resp, err := http.Get("http://www.baidu.com/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // 确保关闭响应体

	// 打印HTTP状态码
	fmt.Printf("resp.StatusCode=%d\n", resp.StatusCode)

	// // 读取响应体
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 打印响应体内容（这里可能会打印很多HTML，所以通常只用于调试）
	// fmt.Println(string(body))
}

// baidu:
// resp.StatusCode=200

// 2024/05/16 18:35:54 Get "http://www.google.com": dial tcp [::1]:80: connect: connection refused
// exit status 1
