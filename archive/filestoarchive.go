package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	// 创建一个打包文件
	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	//使用zip包的创建函数zipWriter写入文件
	zipWriter := zip.NewWriter(outFile)
	//往打包文件中写文件
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test1.txt", "This is a test file"},
		{"test2.txt", "This is another test file"},
		{"test3.txt", "This is yet another test file"},
	}
	//下面将要打包的内容写入打包文件中，依次写入
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	//清理
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
