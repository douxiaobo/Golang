package main

import (
	"net/http"
)

type Refer struct {
	handler http.Handler
	refer   string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if r.Referer() == this.refer {
	// 	this.handler.ServeHTTP(w, r)
	// } else {
	// 	w.WriteHeader(403)
	// }
	this.handler.ServeHTTP(w, r)
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.shirdon.com",
	}
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", referer)
}
