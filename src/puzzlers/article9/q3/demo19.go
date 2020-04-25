package main

import "fmt"

func main() {
	// 键的类型是受限的，而元素却可以是任意类型的
	var m map[string]int

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map: %d\n",
		len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	delete(m, key)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	// 只有 put 操作才会有 nil 问题
	m["two"] = elem // 这里会引发panic。
}
