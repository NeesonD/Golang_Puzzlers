package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	// main 如何接受参数
	flag.StringVar(&name, "name2", "everyone", "The greeting object.")
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	// 解析命令参数
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
