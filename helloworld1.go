package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	// 监听所有网络接口（0.0.0.0）上的 8080 端口
	log.Fatal(http.ListenAndServe(":8080", nil))
}
