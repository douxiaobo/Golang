package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	rangeTemplate := `
	{{if .Chinese}}
	{{range $i, $v:=.Name}}
	{{$i}}=>{{$v}},{{$.Age}},是中国人
	{{end}}
	{{else}}
	{{range .Name}}
	{{.}},{{$.City}},是外国人
	{{end}}
	{{end}}`

	str1 := []string{"小明", "小周"}
	str2 := []string{"汤森", "佩斯"}
	str3 := []string{"小王", "小朱"}
	str4 := []string{"佐罗", "奥丁"}

	type Content struct {
		Name    []string
		Age     int
		City    string
		Chinese bool
	}

	var contents = []Content{
		{str1, 20, "北京", true},
		{str2, 22, "纽约", false},
		{str3, 24, "北京", true},
		{str4, 26, "伦敦", false},
	}
	t := template.Must(template.New("range").Parse(rangeTemplate))
	for _, c := range contents {
		err := t.Execute(os.Stdout, c)
		if err != nil {
			log.Println("executing template failed", err)
		}
	}
}

// 0=>第一次 range,第一次外面的内容

// 1=>有 index 和 value,第一次外面的内容

// 第二次 range,第二次外面的内容

// 没有用 index 和 value,第二次外面的内容
