package main

import "fmt"

// *操作地址对应的内存的值
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func main() {
	first := 100
	second := 200
	// & 传入地址
	swap(&first, &second)
	fmt.Println('1', &first)
	fmt.Println('2', &second)
	fmt.Println('1', first)
	fmt.Println('2', second)
}
