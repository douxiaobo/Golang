package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// dirPath := "/path/to/your/directory"
	dirPath := "./i18n/"
	var filenames []string

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			filenames = append(filenames, filename)
		}
	}

	for _, name := range filenames {
		fmt.Println(name)
	}
}
