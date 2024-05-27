package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Welcome() string {
	return "Welcome"
}

func Doing(name string) string {
	return name + ", growing stronger in Windows"
}

func index(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./index.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	funcs := func() string {
		return "Hello, world!"
	}
	tmpl1, err := template.New("index.html").Funcs(template.FuncMap{"funcs": funcs}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpl1.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	funcMap := template.FuncMap{
		"Welcome": Welcome,
		"Doing":   Doing,
	}
	name := "Microsoft"
	tmpl2, err := template.New("index.html").Funcs(funcMap).Parse("{{Welcome}}\n{{Doing .}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl2.Execute(w, name)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
