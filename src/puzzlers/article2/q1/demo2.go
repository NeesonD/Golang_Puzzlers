package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	// main 如何接受参数
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// 解析命令参数
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
