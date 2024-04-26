package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "a name to say hello to")
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
}

// go build -o hello
// ./hello

// ./hello -name="Go"

// douxiaobo@192 Golang % go run hello_name_command.go
// Hello, world!
// douxiaobo@192 Golang % go run hello_name_command.go -name="Go"
// Hello, Go!
