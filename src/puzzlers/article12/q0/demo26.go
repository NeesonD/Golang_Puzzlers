package main

import "fmt"

// 只要入参和出参相同，那就是函数的一个实现
type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer
	p = printToStd
	p("something")
}
