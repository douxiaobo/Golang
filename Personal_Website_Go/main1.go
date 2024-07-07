package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Website struct {
	Language string
	Title    string
}

var user Website

type LanguageMap map[string]string

var langrange = [...]string{"en", "zh", "es"}

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Header.Get("Accept-Language"))	//zh-CN,zh;q=0.9,en;q=0.8,es;q=0.7
	for _, pref := range strings.Split(r.Header.Get("Accept-Language"), ",") {
		langCode := strings.Split(pref, ";")[0]
		for _, lang := range langrange {
			if langCode == lang {
				user.Language = lang
				goto lang_decide
			}
		}
	}

	if user.Language == "" {
		user.Language = "en"
	}

	// user.Language = strings.Split(r.Header.Get("Accept-Language"), ",")[0][:2]
	// fmt.Fprintf(w, "Detected language: %s\n", user.language)
	// user.Title = "个人主页"
lang_decide:
	// 读取JSON文件
	data, err := ioutil.ReadFile("./public/json/Title.json")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	// 解析JSON到LanguageMap
	var langMap LanguageMap
	err = json.Unmarshal(data, &langMap)
	if err != nil {
		log.Fatalf("error unmarshalling json: %v", err)
	}
	// 根据语言获取Title
	if text, ok := langMap[user.Language]; ok {
		user.Title = text
	} else {
		user.Title = "Homepage"
		log.Println("Language '%s' not found in JSON", user.Language)
	}

	t, err := template.ParseFiles("./public/tmpl/index.html")
	if err != nil {
		fmt.Println("template parsefile failed, error:", err)
		return
	}
	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/", indexHandleFunc)
	http.ListenAndServe(":8080", nil)
}
