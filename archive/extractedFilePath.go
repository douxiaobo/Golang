package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()
	//遍历打包文件中的每一文件/文件夹
	for _, file := range zipReader.Reader.File {
		//打包文件中的文件就像普通的一个文件对象一样
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()
		//指定提取的文件名
		//可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)
		//提取项目或者创建文件夹
		if file.FileInfo().IsDir() {
			//创建文件夹并设置同样的权限
			log.Println("Extracting directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			//提取正常文件
			log.Println("正在提取文件", file.Name)
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()
			//将打包文件中的文件写入到目标文件夹
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
