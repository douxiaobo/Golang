package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// 将路径信息嵌入到HTML中
		htmlResponse := fmt.Sprintf("<html><body><h1>Path: %s</h1></body></html>", path)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(htmlResponse))
	})

	http.ListenAndServe(":80", nil)
}

// Path: /
