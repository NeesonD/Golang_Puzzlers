package main

import (
	"errors"
	"fmt"
)

type oper func(x int, y int) int

func cal(x int, y int, oper2 oper) (int, error) {
	if oper2 == nil {
		return 0, errors.New("函数为空")
	}
	return oper2(x, y), nil
}

type neeson func(x, y int) (int, error)

func getNeeson(oper2 oper) neeson {
	return func(x, y int) (i int, err error) {
		if oper2 == nil {
			return 0, errors.New("nil")
		}
		return oper2(x, y), err
	}
}

func main() {
	f := func(x, y int) int {
		fmt.Println(x)
		fmt.Println(y)
		return x + y
	}

	cal(1, 2, f)
	n := getNeeson(f)
	i, err := n(3, 4)

	fmt.Printf("kkk%d,%v", i, err)

}
