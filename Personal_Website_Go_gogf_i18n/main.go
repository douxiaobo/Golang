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
	Language       string
	Name           string
	Title          string
	NavLinks       []NavLink
	SubNavLinks    []NavLink
	Content        template.HTML
	ContentName    string
	SubContentName string
	FooterLinks    []FooterLink
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

var subcontainlist []string

type ContentNames struct {
	Names []string `json:"names"`
}

type MarathonData struct {
	Zh []MarathonEntry `json:"zh"`
	En []MarathonEntry `json:"en"`
	Es []MarathonEntry `json:"es"`
}

type MarathonEntry struct {
	Id       string `json:"Id"`
	Date     string `json:"Date"`
	City     string `json:"City"`
	Marathon string `json:"Marathon"`
	Project  string `json:"Project"`
}

type TravelEntry struct {
	Id      string `json:"Id"`
	Date    string `json:"Date"`
	Country string `json:"Country"`
}

type TravelData struct {
	Zh []TravelEntry `json:"zh"`
	En []TravelEntry `json:"en"`
	Es []TravelEntry `json:"es"`
}

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/content", http.StripPrefix("/content/", http.FileServer(http.Dir("./content"))))

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
	} else if len(pathParts) == 1 || (len(pathParts) == 2 && pathParts[1] == "") {
		user.Language = pathParts[0]
		user.ContentName = "home"
	} else if len(pathParts) == 2 || (len(pathParts) == 3 && pathParts[2] == "") {
		user.Language = pathParts[0]
		user.ContentName = pathParts[1]
	} else if len(pathParts) == 3 || (len(pathParts) == 4 && pathParts[2] == "") {
		user.Language = pathParts[0]
		user.ContentName = pathParts[1]
		user.SubContentName = pathParts[2]
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

	// 检查内容名称是否有效
	if !isValidContentName(user.ContentName) {
		user.ContentName = "home"
		http.Redirect(w, r, fmt.Sprintf("/%s/%s", user.Language, user.ContentName), http.StatusMovedPermanently)
		return
	}

	// if user.ContentName == "learn" && !isValidSubContentName(user.SubContentName) {
	// 	user.SubContentName = "speak"
	// 	fmt.Println("Invalid sub content name, redirecting to speak")
	// 	http.Redirect(w, r, fmt.Sprintf("/%s/%s", user.Language, user.ContentName), http.StatusMovedPermanently)
	// 	return
	// }

	i18n.SetLanguage(user.Language)
	user.Title = i18n.Translate(ctx, "title")
	user.Name = i18n.Translate(ctx, "name")

	// 根据 ContentName 处理不同的内容
	if user.ContentName == "travel" {
		travelcontent, err := readAndParseTravelContents()
		if err != nil {
			log.Fatalf("Error processing contents: %v", err)
		}
		user.Content = template.HTML(travelcontent)
	} else if user.ContentName == "sports" {
		sportcontent, err := readAndParseSportContents()
		if err != nil {
			log.Fatalf("Error processing sport contents: %v", err)
		}
		user.Content = template.HTML(sportcontent)
	} else if user.ContentName == "learn" {
		if user.SubContentName == "" {
			user.SubContentName = "speak"
		}
		file, err := os.Open(fmt.Sprintf("./content/%s_%s.html", user.SubContentName, user.Language))
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

	} else if (user.ContentName == "home" || user.ContentName == "aboutme" || user.ContentName == "work") && user.Language == "zh" {
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

	if user.ContentName == "learn" {
		user.SubNavLinks, err = readAndParseSubNavLinks()
		if err != nil {
			log.Fatalf("Error processing sub nav links: %v", err)
		}
	}

	if err := tpl.ExecuteTemplate(w, "index.html", user); err != nil {
		log.Println("Template execution error: ", err)
		return
	}
}

func readAndParseTravelContents() (string, error) {
	file, err := os.Open("./travel.json")
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	var travelData TravelData
	err = json.Unmarshal(data, &travelData)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling json: %w", err)
	}

	var content []TravelEntry
	switch user.Language {
	case "zh":
		content = travelData.Zh
	case "en":
		content = travelData.En
	case "es":
		content = travelData.Es
	default:
		content = travelData.En // 默认使用英语
	}

	var result strings.Builder
	var thead bool = true                                                // 表头是否已输出
	var tbody bool = true                                                // 数据行是否已输出
	result.WriteString("<table border='1' style='text-align: center;'>") // 添加表格边框

	for _, entry := range content {
		if thead {
			// 表头
			result.WriteString("<thead><tr>")

			for _, header := range []string{entry.Id, entry.Date, entry.Country} {
				result.WriteString(fmt.Sprintf("<th>%s</th>", header))
			}
			result.WriteString("</tr></thead>")
			thead = false
		} else {
			if tbody {
				// 数据行
				result.WriteString("<tbody>")
				tbody = false
			}

			result.WriteString("<tr>")
			for _, value := range []string{entry.Id, entry.Date, entry.Country} {
				result.WriteString(fmt.Sprintf("<td>%s</td>", value))
			}
			result.WriteString("</tr>")
		}
	}
	result.WriteString("</tbody>")
	result.WriteString("</table>")
	return result.String(), nil
}

func readAndParseSportContents() (string, error) {
	marathon, err := os.Open("./marathon.json")
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	defer marathon.Close()

	data, err := io.ReadAll(marathon)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	var marathonDate MarathonData
	err = json.Unmarshal(data, &marathonDate)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling json: %w", err)
	}

	var content []MarathonEntry
	switch user.Language {
	case "zh":
		content = marathonDate.Zh
	case "en":
		content = marathonDate.En
	case "es":
		content = marathonDate.Es
	default:
		content = marathonDate.En // 默认使用英语
	}

	var result strings.Builder
	var thead bool = true                                                // 表头是否已输出
	var tbody bool = true                                                // 数据行是否已输出
	result.WriteString("<table border='1' style='text-align: center;'>") // 添加表格边框

	for _, entry := range content {
		if thead {
			// 表头
			result.WriteString("<thead><tr>")

			for _, header := range []string{entry.Id, entry.Date, entry.City, entry.Marathon, entry.Project} {
				result.WriteString(fmt.Sprintf("<th>%s</th>", header))
			}
			result.WriteString("</tr></thead>")
			thead = false
		} else {
			if tbody {
				// 数据行
				result.WriteString("<tbody>")
				tbody = false
			}

			result.WriteString("<tr>")
			for _, value := range []string{entry.Id, entry.Date, entry.City, entry.Marathon, entry.Project} {
				result.WriteString(fmt.Sprintf("<td>%s</td>", value))
			}
			result.WriteString("</tr>")
		}
	}
	result.WriteString("</tbody>")
	result.WriteString("</table>")
	return result.String(), nil
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

// func isValidSubContentName(contentName string) bool {
// 	for _, c := range subcontainlist {
// 		if contentName == c {
// 			return true
// 		}
// 	}
// 	return false
// }

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
	var navlink NavLink
	for _, name := range containlist {
		navlink = NavLink{Link_name: i18n.Translate(ctx, name), Link_url: fmt.Sprintf("/%s/%s", user.Language, name)}
		navlinks = append(navlinks, navlink)
	}
	return navlinks, nil
}

func readAndParseSubNavLinks() ([]NavLink, error) {
	var subnavlinks []NavLink

	contentNames, err := readContentNamesFile("./sub_" + user.ContentName + ".json")
	if err != nil {
		log.Fatalf("Error reading content names file: %v", err)
	}
	subcontainlist = contentNames.Names

	for _, name := range subcontainlist {
		subnavlink := NavLink{Link_name: i18n.Translate(ctx, name), Link_url: fmt.Sprintf("/%s/%s/%s", user.Language, user.ContentName, name)}
		subnavlinks = append(subnavlinks, subnavlink)
	}
	return subnavlinks, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in main: %v", r)
		}
	}()

	log.Println("Starting HTTP server...")
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
