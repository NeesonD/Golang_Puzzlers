package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	// 开箱即用，延迟初始化
	var list list.List
	list.PushBack("1")
	list.PushBack("1")
	fmt.Println(list.Len())

	var r ring.Ring
	r2 := ring.New(5)
	r.Link(r2)
	fmt.Println(r.Len())
}
