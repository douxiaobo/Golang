package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	//打开一个gzip文件
	gzipFile, err := os.Open("output.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer gzipFile.Close()
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer gzipReader.Close()
	//解压缩到一个writer，它是一个file write
	outfileWriter, err := os.Create("output.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer outfileWriter.Close()
	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatalln(err)
	}
}
