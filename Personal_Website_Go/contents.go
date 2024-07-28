package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Contents struct {
	Zh string `json:"zh"`
	En string `json:"en"`
	Es string `json:"es"`
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

func readAndParseTravelContents(fileName string) (TravelData, error) {
	file, err := os.Open(fmt.Sprintf("./public/json/%s.json", fileName))
	if err != nil {
		return TravelData{}, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return TravelData{}, fmt.Errorf("error reading file: %w", err)
	}

	var travelData TravelData
	err = json.Unmarshal(data, &travelData)
	if err != nil {
		return TravelData{}, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return travelData, nil
}

func readAndParseWorkContents() (WorkContent, error) {
	workFile, err := os.Open("./public/json/work.yaml")
	if err != nil {
		return WorkContent{}, fmt.Errorf("error opening file: %w", err)
	}
	defer workFile.Close()

	var workContent WorkContent
	decoder := yaml.NewDecoder(workFile)
	err = decoder.Decode(&workContent)
	if err != nil {
		return WorkContent{}, fmt.Errorf("error decoding yaml: %w", err)
	}
	return workContent, nil
}

func readAndParseSportContents() (string, error) {
	marathon, err := os.Open("./public/json/marathon.json")
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
	result.WriteString("<table border='1'>") // 添加表格边框

	// 表头
	result.WriteString("<thead><tr>")
	for _, header := range []string{"Id", "Date", "City", "Marathon", "Project"} {
		result.WriteString(fmt.Sprintf("<th>%s</th>", header))
	}
	result.WriteString("</tr></thead>")

	// 数据行
	result.WriteString("<tbody>")
	for _, entry := range content {
		result.WriteString("<tr>")
		for _, value := range []string{entry.Id, entry.Date, entry.City, entry.Marathon, entry.Project} {
			result.WriteString(fmt.Sprintf("<td>%s</td>", value))
		}
		result.WriteString("</tr>")
	}
	result.WriteString("</tbody>")

	result.WriteString("</table>")

	return result.String(), nil

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

func getTravelContents(traveldata TravelData, language string) []TravelEntry {
	var content []TravelEntry
	switch language {
	case "zh":
		content = traveldata.Zh
	case "en":
		content = traveldata.En
	case "es":
		content = traveldata.Es
	default:
		content = traveldata.En // Default to English
	}

	return content
}

func getWorkContent(workcontent WorkContent, language string) string {
	switch language {
	case "zh":
		return workcontent.Zh
	case "en":
		return workcontent.En
	case "es":
		return workcontent.Es
	default:
		return "Work Content"
	}
}
