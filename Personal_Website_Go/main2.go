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
	Language    string
	Title       string
	Menu        string
	FooterLinks []FooterLink
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

// type Footers struct {
// 	Zh []string `json:"zh"`
// 	En []string `json:"en"`
// 	Es []string `json:"es"`
// }

type FooterLink struct {
	Code string `json:"code"` // 假设我们使用"code"作为JSON中的键，代表语言代码
	Name string `json:"name"` // 假设我们使用"name"作为JSON中的键，代表语言名称
}

type FooterLinks []FooterLink

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	// 从URL路径中获取语言后缀
	path := strings.TrimPrefix(r.URL.Path, "/")
	lang := strings.Split(path, "/")[0]

	// 如果URL中没有语言后缀，并且请求的是根路径，从Accept-Language头部获取
	if lang == "" {
		lang = getLanguageFromHeader(r.Header.Get("Accept-Language"))

		// 如果找到了匹配的语言，执行重定向
		if lang != "" {
			http.Redirect(w, r, "/"+lang, http.StatusFound)
			return
		}
	}

	// 检查语言是否在支持的范围内
	found := false
	for _, supportedLang := range langrange {
		if lang == supportedLang {
			user.Language = lang
			found = true
			break
		}
	}

	// 如果没有找到匹配的语言，默认使用英语
	if !found {
		user.Language = "en"
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

	// 读取Footer.json文件
	footerData, err := ioutil.ReadFile("./public/json/Footer.json")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	// 解析JSON到新的Footers结构体
	var footerJSON struct {
		Links []FooterLink `json:"links"`
	}
	err = json.Unmarshal(footerData, &footerJSON)
	if err != nil {
		log.Fatalf("error unmarshalling json: %v", err)
	}
	// 将解析后的链接设置到user结构体中
	user.FooterLinks = footerJSON.Links

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

// func getFooter(footers Footers, language string) string {
// 	switch language {
// 	case "zh":
// 		return strings.Join(footers.Zh, ", ")
// 	case "en":
// 		return strings.Join(footers.En, ", ")
// 	case "es":
// 		return strings.Join(footers.Es, ", ")
// 	default:
// 		return "Footer"
// 	}
// }

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
