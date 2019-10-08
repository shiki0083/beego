package main

import "fmt"

func test() {
	//捕获异常 不影响其他进行
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err", err)
		}
	}()

	b := 0
	a := 100 / b
	fmt.Println("hhhh", a)
	return
}

func main() {

	test()
	var a []int
	a = append(a, a...)
	fmt.Println("a", a)
}
