package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gctx"
)

type User struct {
	Language    string
	Name        string
	Title       string
	NavLinks    []NavLink
	Content     template.HTML
	ContentName string
	FooterLinks []FooterLink
}

type FooterLink struct {
	Footer_language_short string
	Footer_language_long  string
}

type NavLink struct {
	Link_name string
	Link_url  string
}

var user User

var err error

type FooterLinks []FooterLink

var (
	ctx  = gctx.New()
	i18n = gi18n.New()
)

var languages_ranges []string

var containlist []string

type ContentNames struct {
	Names []string `json:"names"`
}

func init() {
	dirPath := "./i18n/"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			languages_ranges = append(languages_ranges, filename)
		}
	}

	contentNames, err := readContentNamesFile("./content_names.json")
	if err != nil {
		log.Fatalf("Error reading content names file: %v", err)
	}
	containlist = contentNames.Names

	// for _, name := range containlist {
	// 	fmt.Println(name)
	// }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	// 从URL路径中获取语言后缀
	path := strings.TrimPrefix(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")

	// 如果路径为空或仅为 "/"
	if len(pathParts) == 0 || (len(pathParts) == 1 && pathParts[0] == "") {
		// 设置默认语言和内容名称
		user.Language = getLanguageFromHeader(r.Header.Get("Accept-Language"))
		user.ContentName = "home"
	} else if len(pathParts) == 1 || len(pathParts) == 2 && pathParts[1] == "" {
		user.Language = pathParts[0]
		user.ContentName = "home"
	} else if len(pathParts) == 2 {
		user.Language = pathParts[0]
		user.ContentName = pathParts[1]
	} else {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// 检查语言是否在支持的范围内
	if !isValidLang(user.Language) {
		user.Language = "en"
		http.Redirect(w, r, fmt.Sprintf("/%s/%s", user.Language, user.ContentName), http.StatusMovedPermanently)
		return
	}

	if !isValidContentName(user.ContentName) {
		user.ContentName = "home"
		http.Redirect(w, r, fmt.Sprintf("/%s/%s", user.Language, user.ContentName), http.StatusMovedPermanently)
		return
	}

	i18n.SetLanguage(user.Language)
	user.Title = i18n.Translate(ctx, "title")
	user.Name = i18n.Translate(ctx, "name")

	if user.ContentName == "aboutme" && user.Language == "zh" {
		file, err := os.Open(fmt.Sprintf("./content/%s_%s.html", user.ContentName, user.Language))
		if err != nil {
			log.Printf("Error opening content file: %v", err)
			return
		}
		defer file.Close()
		// 读取文件内容
		contentBytes, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Error reading content file: %v", err)
			return // 同样处理错误
		}

		// 将内容转换为 template.HTML 类型
		user.Content = template.HTML(contentBytes)
	} else {
		content := "<p>This is " + user.ContentName + "</p>"
		user.Content = template.HTML(content)
	}

	user.FooterLinks, err = readAndParseFooterLinks()
	if err != nil {
		log.Fatalf("Error processing footer links: %v", err)
	}

	user.NavLinks, err = readAndParseNavLinks()
	if err != nil {
		log.Fatalf("Error processing nav links: %v", err)
	}

	t, err := template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Template parsing error:", err)
		return
	}
	if err := t.ExecuteTemplate(w, "index.html", user); err != nil {
		log.Println("Template execution error: ", err)
		return
	}
}

func getLanguageFromHeader(header string) string {
	langs := strings.Split(header, ",")
	for _, lang := range langs {
		trimmedLang := strings.TrimSpace(lang)
		for _, supportedLang := range languages_ranges {
			if strings.HasPrefix(trimmedLang, supportedLang) {
				return supportedLang
			}
		}
	}
	return "en"
}

func isValidLang(lang string) bool {
	for _, l := range languages_ranges {
		if lang == l {
			return true
		}
	}
	return false
}

func isValidContentName(contentName string) bool {
	for _, c := range containlist {
		if contentName == c {
			return true
		}
	}
	return false
}

func readAndParseFooterLinks() (FooterLinks, error) {
	var footerlinks []FooterLink
	var footerlink FooterLink
	for _, lang := range languages_ranges {
		footerlink.Footer_language_short = lang
		i18n.SetLanguage(lang)
		footerlink.Footer_language_long = i18n.Translate(ctx, "language_full")
		footerlinks = append(footerlinks, footerlink)
	}
	i18n.SetLanguage(user.Language)

	return footerlinks, nil
}

func readContentNamesFile(filePath string) (*ContentNames, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var contentNames ContentNames
	err = json.NewDecoder(file).Decode(&contentNames)
	if err != nil {
		return nil, err
	}

	return &contentNames, nil
}

func readAndParseNavLinks() ([]NavLink, error) {
	var navlinks []NavLink
	for _, name := range containlist {
		navlink := NavLink{Link_name: i18n.Translate(ctx, name), Link_url: fmt.Sprintf("/%s/%s", user.Language, name)}
		navlinks = append(navlinks, navlink)
	}
	return navlinks, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in main: %v", r)
		}
	}()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Starting HTTP server...")
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
