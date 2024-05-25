package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	name := "Microsoft"
	Title := "Welcome to Microsoft"
	Content := "This is a sample website for Microsoft"
	t.Execute(w, map[string]interface{}{
		"Name":    name,
		"Title":   Title,
		"Content": Content,
		"Footer":  "Copyright © 2021 Microsoft",
	})
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
