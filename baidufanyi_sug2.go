package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type translateHandler struct{}

func (h translateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Path[1:] // 获取URL路径中"/"后的部分作为word
	// 在这里添加翻译逻辑，例如调用外部API或数据库查询
	fmt.Fprintf(w, "You requested to translate: %s\n", word)
	jsonData := baidufanyi(word)
	if jsonData != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// w.Write(jsonData)
		fmt.Fprintf(w, "\nTranslation suggestion: %s\n", jsonData)
	} else {
		http.Error(w, "Failed to fetch translation suggestions", http.StatusInternalServerError)
	}
	var userWord Response
	err := json.Unmarshal(jsonData, &userWord)
	if err != nil {
		http.Error(w, "Failed to parse translation suggestions", http.StatusInternalServerError)
		return
	}
	var result = string(userWord.Data[0].V)
	fmt.Fprintf(w, "Translation suggestion: %s\n", result)
	fmt.Println("Translation suggestion:", result)
}

func main() {
	server := &http.Server{
		// Addr: "127.0.0.1:8080",
		Addr: "localhost:8080",
	}
	var userWord string
	fmt.Printf("Please input the word you want to translate: ")
	fmt.Scan(&userWord)

	// 使用用户输入的单词创建实际的路由
	http.HandleFunc("/"+userWord, translateHandler{}.ServeHTTP)
	fmt.Println("Server is listening on :8080")
	fmt.Println("Please visit http://localhost:8080/" + userWord + " to see the translation suggestion.")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}

type Response struct {
	Errno int         `json:"errno"`
	Data  []DataEntry `json:"data"`
	Logid int64       `json:"logid"`
}

type DataEntry struct {
	K string `json:"k"`
	V string `json:"v"`
}

func baidufanyi(word string) []byte {
	// 调用百度翻译API，将word翻译成中文
	url := "https://fanyi.baidu.com/sug"
	method := "POST"

	payload := strings.NewReader("kw=" + word)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "BAIDUID=4C26A499B77432870FC4DC861A9DEA01:FG=1; BAIDUID_BFESS=4C26A499B77432870FC4DC861A9DEA01:FG=1")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// fmt.Println(string(body))

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Output the decoded JSON data
	jsonData, _ := json.MarshalIndent(response, "", "  ")
	// fmt.Println(string(jsonData))
	return jsonData
}
