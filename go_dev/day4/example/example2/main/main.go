package main

import "fmt"

func main()  {
	//声明一个10长度的数组
	var a[10]int
	a[0] = 100
	a[1] = 1

	arr2 := a
	arr2[0] = 222
	fmt.Println("aaaa",a)
	fmt.Println("aaaa",arr2)

}