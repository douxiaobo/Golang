package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("请输入一个名字")
		return
	}
	fmt.Println("hello", strings.ToUpper(args[1]))
}
