// title.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Titles 结构体定义
type Titles struct {
	Zh string `json:"zh"`
	En string `json:"en"`
	Es string `json:"es"`
}

// readAndParseTitles 读取并解析Title.json
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

// getTitle 根据语言获取Title
func getTitle(titles Titles, language string) string {
	switch language {
	case "zh":
		return titles.Zh
	case "en":
		return titles.En
	case "es":
		return titles.Es
	default:
		return titles.En // 默认使用英语
	}
}
