package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Website struct {
	Language    string
	Title       string
	Menu        string
	Content     string
	ContentName string
	FooterLinks []FooterLink
	CookieName  string
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

type Contents struct {
	Zh string `json:"zh"`
	En string `json:"en"`
	Es string `json:"es"`
}

type FooterLink struct {
	Code string `json:"code"` // 假设我们使用"code"作为JSON中的键，代表语言代码
	Name string `json:"name"` // 假设我们使用"name"作为JSON中的键，代表语言名称
}

type FooterLinks []FooterLink

var content_name string

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	{
		// 从URL路径中获取语言后缀
		path := strings.TrimPrefix(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")

		// 初始化Website结构体的ContentName字段
		user.ContentName = "home" // 默认值

		// 如果路径中有足够的部分，那么第一部分是语言，第二部分是内容名称
		if len(pathParts) >= 2 {
			user.Language = pathParts[0]
			user.ContentName = pathParts[1]
		}

		// 尝试从cookie中获取语言设置
		cookie, err := r.Cookie(user.CookieName)
		if err == nil && cookie.Value != "" {
			user.Language = cookie.Value
		} else if user.Language == "" {
			// 如果cookie不存在或者URL中没有语言后缀，从Accept-Language头部获取
			user.Language = getLanguageFromHeader(r.Header.Get("Accept-Language"))
		}

		// 检查语言是否在支持的范围内
		found := false
		for _, supportedLang := range langrange {
			if user.Language == supportedLang {
				found = true
				break
			}
		}

		// 如果没有找到匹配的语言，默认使用英语
		if !found {
			user.Language = "en"
		}

		// 设置cookie来记住用户的语言选择
		http.SetCookie(w, &http.Cookie{
			Name:  user.CookieName,
			Value: user.Language,
			Path:  "/",
		})

		if content_name == "" {
			content_name = "home" // 这里替换为你的默认内容名称
		}

		// 如果URL中没有语言后缀或内容名称，重定向到带有语言后缀和内容名称的URL
		if user.Language == "" || user.ContentName == "" {
			http.Redirect(w, r, "/"+user.Language+"/"+user.ContentName, http.StatusFound)
			return
		}
	}

	// 使用新函数读取并解析Title.json
	titles, err := readAndParseTitles()
	if err != nil {
		log.Fatalf("Error processing titles: %v", err)
	}

	// 读取Menu.json文件
	menus, err := readAndParseMenus()
	if err != nil {
		log.Fatalf("Error processing menus: %v", err)
	}

	// 读取内容的JSON文件
	contents, err := readAndParseContents(content_name)
	if err != nil {
		log.Fatalf("Error processing contents: %v", err)
	}

	user.FooterLinks, err = readAndParseFooterLinks()
	if err != nil {
		log.Fatal("Error processing footer links: %v", err)
	}

	// 根据语言获取Title
	user.Title = getTitle(titles, user.Language)

	// 根据语言获取Menu
	user.Menu = getMenu(menus, user.Language)

	user.Content = getContents(contents, user.Language)

	// fmt.Printf("Footer Links: %+v\n", user.FooterLinks)
	// fmt.Printf("Content Name: %s\n", user.ContentName)

	t, err := template.ParseFiles("./public/tmpl/index1.html")
	if err != nil {
		fmt.Println("template parsefile failed, error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, map[string]interface{}{
		"Title":       user.Title,
		"Menu":        user.Menu,
		"Content":     user.Content,
		"ContentName": user.ContentName,
		"User":        user,
		"FooterLinks": user.FooterLinks,
	})
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func readAndParseTitles() (Titles, error) {
	titleFile, err := os.Open("./public/json/Title.json")
	if err != nil {
		return Titles{}, fmt.Errorf("error opening file: %w", err)
	}
	defer titleFile.Close()

	titleData, err := io.ReadAll(titleFile)
	if err != nil {
		return Titles{}, fmt.Errorf("error reading file: %w", err)
	}

	var titles Titles
	err = json.Unmarshal(titleData, &titles)
	if err != nil {
		return Titles{}, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return titles, nil
}

func readAndParseMenus() (Menus, error) {
	menuFile, err := os.Open("./public/json/Menu1.json")
	if err != nil {
		return Menus{}, fmt.Errorf("error opening file: %w", err)
	}
	defer menuFile.Close()

	menuData, err := io.ReadAll(menuFile)
	if err != nil {
		return Menus{}, fmt.Errorf("error reading file: %w", err)
	}

	var menus Menus
	err = json.Unmarshal(menuData, &menus)
	if err != nil {
		return Menus{}, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return menus, nil
}

func readAndParseContents(name string) (Contents, error) {
	contentFile, err := os.Open(fmt.Sprintf("./public/json/%s.json", name))
	if err != nil {
		return Contents{}, fmt.Errorf("error opening file: %w", err)
	}
	defer contentFile.Close()

	contentData, err := io.ReadAll(contentFile)
	if err != nil {
		return Contents{}, fmt.Errorf("error reading file: %w", err)
	}

	var contents Contents
	err = json.Unmarshal(contentData, &contents)
	if err != nil {
		return Contents{}, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return contents, nil
}

func readAndParseFooterLinks() (FooterLinks, error) {
	// 读取Footer.json文件
	footerData, err := ioutil.ReadFile("./public/json/Footer.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// 解析JSON到新的Footers结构体
	var footerJSON struct {
		Links []FooterLink `json:"links"`
	}
	err = json.Unmarshal(footerData, &footerJSON)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}
	// 返回解析后的链接
	return footerJSON.Links, nil
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

func getContents(contents Contents, language string) string {
	switch language {
	case "zh":
		return contents.Zh
	case "en":
		return contents.En
	case "es":
		return contents.Es
	default:
		return "Content"
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
	user = Website{
		CookieName: "preferred_language",
	}
	http.HandleFunc("/", HandleFunc)
	http.ListenAndServe(":8080", nil)
}