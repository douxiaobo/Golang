package main

//这个代码失败，作废

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
	// MenuLinks   []MenuLink
	Content     string
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
// type Menus struct {
// 	Zh []MenuLink `json:"zh"`
// 	En []MenuLink `json:"en"`
// 	Es []MenuLink `json:"es"`
// }

type Contents struct {
	Zh string `json:"zh"`
	En string `json:"en"`
	Es string `json:"es"`
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

// type MenuLink struct {
// 	Name string `json:"name"`
// 	link string `json:"link"`
// }

var content_name string

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	{
		var lang string

		// 从URL路径中获取语言和内容名称
		path := strings.TrimPrefix(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")

		// 检查URL是否有足够的部分来解析语言和内容名称
		if len(pathParts) >= 2 {
			lang = pathParts[0]
			content_name = pathParts[1]
		} else {
			// 如果URL中没有内容名称，从Accept-Language头部获取语言
			lang = getLanguageFromHeader(r.Header.Get("Accept-Language"))
			content_name = "home" // 默认内容名称

			// 如果找到了匹配的语言，执行重定向
			if lang != "" {
				http.Redirect(w, r, "/"+lang+"/"+content_name, http.StatusFound)
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
	}

	{
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
	}

	// {
	// 	// 读取Menu.json文件
	// 	menuData, err := ioutil.ReadFile("./public/json/Menu.json")
	// 	if err != nil {
	// 		log.Fatalf("error reading file: %v", err)
	// 	}
	// 	// 解析JSON到Menus结构体
	// 	var menus Menus
	// 	err = json.Unmarshal(menuData, &menus)
	// 	if err != nil {
	// 		log.Fatalf("error unmarshalling json: %v", err)
	// 	}
	// 	// 根据语言获取Menu
	// 	user.MenuLinks = getMenuLinks(menus, user.Language)
	// }

	{
		contentData, err := ioutil.ReadFile("./public/json/" + content_name + ".json")
		if err != nil {
			log.Fatal("error reading file: %v", err)
		}
		var contents Contents
		err = json.Unmarshal(contentData, &contents)
		if err != nil {
			log.Fatalf("error unmarshalling json: %v", err)
		}
		user.Content = getContents(contents, user.Language)
	}

	{
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
	}

	{
		t, err := template.ParseFiles("./public/tmpl/index.html")
		if err != nil {
			fmt.Println("template parsefile failed, error:", err)
			return
		}
		t.Execute(w, user)
	}
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
		return "Title"
	}
}

// // 获取菜单
// func getMenuLinks(menus Menus, language string) []MenuLink {
// 	switch language {
// 	case "zh":
// 		return menus.Zh
// 	case "en":
// 		return menus.En
// 	case "es":
// 		return menus.Es
// 	default:
// 		return []MenuLink{} // 返回空切片
// 	}
// }

func getContents(contents Contents, language string) string {
	switch language {
	case "zh":
		return contents.Zh
	case "en":
		return contents.En
	case "es":
		return contents.Es
	default:
		return "Home"
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
