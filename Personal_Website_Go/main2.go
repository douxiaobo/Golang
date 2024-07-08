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
	Menu     string
}

var user Website

type LanguageMap map[string]string

var langrange = [...]string{"en", "zh", "es"}

// 定义一个结构体来匹配你的JSON数据结构
type Titles struct {
	Zh string `json:"zh"`
	En string `json:"en"`
	Es string `json:"es"`
}

// 同样的结构体用于Menu.json
type Menus struct {
	Zh []string `json:"zh"`
	En []string `json:"en"`
	Es []string `json:"es"`
}

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	// 从URL路径中获取语言后缀
	path := strings.TrimPrefix(r.URL.Path, "/")
	lang := strings.Split(path, "/")[0]

	// 检查语言是否在支持的范围内
	found := false
	for _, supportedLang := range langrange {
		if lang == supportedLang {
			user.Language = lang
			found = true
			break
		}
	}

	// 如果URL中没有语言后缀，从Accept-Language头部获取
	if !found {
		lang = getLanguageFromHeader(r.Header.Get("Accept-Language"))
		if lang != "" {
			user.Language = lang
		} else {
			user.Language = "en" // 如果没有找到匹配的语言，默认使用英语
		}
	}

	// 读取Title.json文件
	titleData, err := ioutil.ReadFile("./public/json/Title.json")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	// 解析JSON到Titles结构体
	var titles Titles
	err = json.Unmarshal(titleData, &titles)
	if err != nil {
		log.Fatalf("error unmarshalling json: %v", err)
	}
	// 根据语言获取Title
	user.Title = getTitle(titles, user.Language)

	// 读取Menu.json文件
	menuData, err := ioutil.ReadFile("./public/json/Menu.json")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	// 解析JSON到Menus结构体
	var menus Menus
	err = json.Unmarshal(menuData, &menus)
	if err != nil {
		log.Fatalf("error unmarshalling json: %v", err)
	}
	// 根据语言获取Menu
	user.Menu = getMenu(menus, user.Language)

	// // 读取JSON文件
	// data, err := ioutil.ReadFile("./public/json/Title.json")
	// if err != nil {
	// 	log.Fatalf("error reading file: %v", err)
	// }
	// // 解析JSON到LanguageMap
	// var langMap LanguageMap
	// err = json.Unmarshal(data, &langMap)
	// if err != nil {
	// 	log.Fatalf("error unmarshalling json: %v", err)
	// }
	// // 根据语言获取Title
	// if text, ok := langMap[user.Language]; ok {
	// 	user.Title = text
	// } else {
	// 	user.Title = "Homepage"
	// 	log.Println("Language '%s' not found in JSON", user.Language)
	// }

	// data, err = ioutil.ReadFile("./public/json/Menu.json")
	// if err != nil {
	// 	log.Fatal("error reading file: %v", err)
	// }
	// err = json.Unmarshal(data, &langMap)
	// if err != nil {
	// 	log.Fatal("error unmarshalling json: %v", err)
	// }
	// if text, ok := langMap[user.Language]; ok {
	// 	user.Menu = text
	// } else {
	// 	user.Menu = "Menu"
	// 	log.Println("Language '%s' not found in JSON", user.Language)
	// }

	t, err := template.ParseFiles("./public/tmpl/index.html")
	if err != nil {
		fmt.Println("template parsefile failed, error:", err)
		return
	}
	t.Execute(w, user)
}

// 获取标题
func getTitle(titles Titles, language string) string {
	switch language {
	case "zh":
		return titles.Zh
	case "en":
		return titles.En
	case "es":
		return titles.Es
	default:
		return "Homepage"
	}
}

// 获取菜单
func getMenu(menus Menus, language string) string {
	switch language {
	case "zh":
		return strings.Join(menus.Zh, ", ")
	case "en":
		return strings.Join(menus.En, ", ")
	case "es":
		return strings.Join(menus.Es, ", ")
	default:
		return "Menu"
	}
}

// 解析Accept-Language头部并返回最优先的语言
func getLanguageFromHeader(header string) string {
	// 这里可以实现对Accept-Language头部的解析逻辑
	// 注意，这可能涉及到质量因子(q-factor)的处理
	// 为简化起见，这里仅返回第一个语言代码
	parts := strings.Split(header, ",")
	if len(parts) > 0 {
		lang := strings.Split(parts[0], ";")[0]
		for _, supportedLang := range langrange {
			if lang == supportedLang {
				return lang
			}
		}
	}
	return ""
}

func main() {
	http.HandleFunc("/", indexHandleFunc)
	http.ListenAndServe(":8080", nil)
}
