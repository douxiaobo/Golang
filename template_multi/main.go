package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func tmplSample(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.html", "./ul.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	user := UserInfo{
		Name:   "Douxiaobao",
		Gender: "Male",
		Age:    25,
	}
	err = tmpl.Execute(w, user)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", tmplSample)
	http.ListenAndServe(":8080", nil)
}
