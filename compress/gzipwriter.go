package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	outputFile, err := os.Create("output.gz")
	if err != nil {
		log.Fatalln(err)
	}
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()
	//当写入gzip writer数据时，它会依次压缩数据并写入底层的文件中
	_, err = gzipWriter.Write([]byte("Hello, world!"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Data written to output.gz")
}
