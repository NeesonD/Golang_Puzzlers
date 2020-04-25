package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	var count int32
	trigger := func(i int32, f func()) {
		for {
			if n := atomic.LoadInt32(&count); n == i {
				f()
				atomic.AddInt32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for u := int32(0); u <= 10; u++ {
		go func(i int32) {
			f := func() {
				fmt.Print(i)
			}
			trigger(i, f)
		}(u)
	}

	trigger(10, func() {})

}
