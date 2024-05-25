package main

import (
	"net/http"
)

type Refer struct {
	handler http.Handler
	refer   string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the homepage!"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.shirdon.com",
	}

	http.HandleFunc("/", rootHandler) // 添加根路径的处理函数
	http.HandleFunc("/hello", hello)

	// 使用 referer 中间件包装整个服务器
	http.ListenAndServe(":8080", referer)
}
