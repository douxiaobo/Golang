package main

import (
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8080", nil)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := UserInfo{
		Name:   "Douxiaobing",
		Gender: "Male",
		Age:    25,
	}
	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
