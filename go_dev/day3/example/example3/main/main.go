package main

import (
	"fmt"
)

func add(a int, arg ...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}(a, b)

	return result
}

func main() {
	sum := add(10, 3, 3)
	fmt.Println(sum)

	var i int = 0
	defer fmt.Println("aaa", i)
	defer fmt.Println("second", i)

	i = 10
	fmt.Println("ccc", i)
	fmt.Println("result", test(100, 200))
}
