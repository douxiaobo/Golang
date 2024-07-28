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

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	{
		// 从URL路径中获取语言后缀
		path := strings.TrimPrefix(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")

		// 如果路径中有足够的部分，那么第一部分是语言，第二部分是内容名称
		if len(pathParts) >= 2 {
			user.Language = pathParts[0]
			user.ContentName = pathParts[1]
			if !(isValidLang(pathParts[0]) && isValidContentName(pathParts[1])) {
				http.Error(w, "Invaid URL", http.StatusBadGateway)
				return
			}
		} else {
			// 尝试从cookie中获取语言设置
			cookie, err := r.Cookie(user.CookieName)
			if err == nil && cookie.Value != "" {
				user.Language = cookie.Value
			} else if user.Language == "" {
				// 如果cookie不存在或者URL中没有语言后缀，从Accept-Language头部获取
				user.Language = getLanguageFromHeader(r.Header.Get("Accept-Language"))

			}
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

		if user.ContentName == "" {
			user.ContentName = "home" // 这里替换为你的默认内容名称
		}

	}

	// 使用新函数读取并解析Title.json
	titles, err := readAndParseTitles()
	if err != nil {
		log.Fatalf("Error processing titles: %v", err)
	}

	// 根据语言获取Title
	user.Title = getTitle(titles, user.Language)

	// 读取Menu.json文件
	menus, err := readAndParseMenus()
	if err != nil {
		log.Fatalf("Error processing menus: %v", err)
	}

	// 根据语言获取Menu
	user.Menu = getMenu(menus, user.Language)

	if user.ContentName == "travel" {
		travelcontent, err := readAndParseTravelContents("travel")
		if err != nil {
			log.Fatalf("Error processing contents: %v", err)
		}
		user.Travel = getTravelContents(travelcontent, user.Language)
	} else if user.ContentName == "work" {
		workcontent, err := readAndParseWorkContents()
		if err != nil {
			log.Fatalf("Error processing work contents: %v", err)
		}
		user.Content = getWorkContent(workcontent, user.Language)
	} else if user.ContentName == "sport" {
		sportcontent, err := readAndParseSportContents()
		if err != nil {
			log.Fatalf("Error processing sport contents: %v", err)
		}
		user.Content = sportcontent
	} else {
		// 读取内容的JSON文件
		contents, err := readAndParseContents(user.ContentName)
		if err != nil {
			log.Fatalf("Error processing contents: %v", err)
		}
		user.Content = getContents(contents, user.Language)

	}

	user.FooterLinks, err = readAndParseFooterLinks()
	if err != nil {
		log.Fatal("Error processing footer links: %v", err)
	}

	t, err := template.ParseFiles("./public/tmpl/index1.html")
	if err != nil {
		fmt.Println("template parsefile failed, error:", err)
		http.Error(w, "Internal Server Error1", http.StatusInternalServerError)
		return
	}

	// log.Println("User Title before rendering:", user.Title)

	if err = t.Execute(w, user); err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Internal Server Error2", http.StatusInternalServerError)
		return
	}
}

func isValidLang(lang string) bool {
	for _, l := range langrange {
		if lang == l {
			return true
		}
	}
	return false
}

func isValidContentName(contentName string) bool {
	for _, c := range containtlist {
		if contentName == c {
			return true
		}
	}
	return false
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

// 解析Accept-Language头部并返回最优先的语言
func getLanguageFromHeader(header string) string {
	langs := strings.Split(header, ",")
	for _, lang := range langs {
		trimmedLang := strings.TrimSpace(lang)
		for _, supportedLang := range langrange {
			if strings.HasPrefix(trimmedLang, supportedLang) {
				return supportedLang
			}
		}
	}
	return ""
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in main: %v", r)
		}
	}()
	user = Website{
		CookieName: "preferred_language",
	}
	log.Println("Starting HTTP server...")
	http.HandleFunc("/", HandleFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

// douxiaobo@192 Personal_Website_Go % go run main3.go
// 2024/07/18 15:56:51 Starting HTTP server...
// 2024/07/18 15:56:51 ListenAndServe: listen tcp :8080: bind: address already in use
// exit status 1
// douxiaobo@192 Personal_Website_Go % lsof -i :8080
// COMMAND  PID      USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
// main3   4960 douxiaobo    3u  IPv6 0x624d3e8b6cf3a7e8      0t0  TCP *:http-alt (LISTEN)
// main3   4960 douxiaobo    7u  IPv6 0xb17d9f4521718fb1      0t0  TCP localhost:http-alt->localhost:56382 (CLOSED)
// main3   4960 douxiaobo    8u  IPv6 0xe8d415cc13afb679      0t0  TCP localhost:http-alt->localhost:56383 (CLOSED)
// douxiaobo@192 Personal_Website_Go % kill 4960
// douxiaobo@192 Personal_Website_Go %

// douxiaobo@192 Personal_Website_Go % go run main3.go
// 2024/07/18 16:04:16 Starting HTTP server...
// 2024/07/18 16:04:16 ListenAndServe: listen tcp :8080: bind: address already in use
// exit status 1
// douxiaobo@192 Personal_Website_Go % lsof -i :8080
// COMMAND  PID      USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
// main3   4960 douxiaobo    3u  IPv6 0x624d3e8b6cf3a7e8      0t0  TCP *:http-alt (LISTEN)
// main3   4960 douxiaobo    7u  IPv6 0xb17d9f4521718fb1      0t0  TCP localhost:http-alt->localhost:56382 (CLOSED)
// main3   4960 douxiaobo    8u  IPv6 0xe8d415cc13afb679      0t0  TCP localhost:http-alt->localhost:56383 (CLOSED)
// douxiaobo@192 Personal_Website_Go % kill -9 4960

// # 初始化模块
// go mod init github.com/douxiaobo/Personal_Website_Go

// # 添加 yaml.v3 的依赖
// go get gopkg.in/yaml.v3

// # 确保依赖是最新的
// go mod tidy

// # 构建并运行你的程序
// go run .

// douxiaobo@192 Personal_Website_Go % go run main3.go
// main3.go:14:2: no required module provides package gopkg.in/yaml.v3: go.mod file not found in current directory or any parent directory; see 'go help modules'
// douxiaobo@192 Personal_Website_Go % go mod init personal_website
// go: creating new go.mod: module personal_website
// go: to add module requirements and sums:
// 	go mod tidy
// douxiaobo@192 Personal_Website_Go % go get gopkg.in/yaml.v3
// go: added gopkg.in/yaml.v3 v3.0.1
// douxiaobo@192 Personal_Website_Go % go mod tidy
// go: downloading gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
// douxiaobo@192 Personal_Website_Go % go run main3.go
// 2024/07/21 15:38:16 Starting HTTP server...
