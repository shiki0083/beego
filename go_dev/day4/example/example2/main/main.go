package main

import "fmt"

func test() {
	var arr [5]int
	for key, value := range arr {
		fmt.Println("i", key, "is", arr[value]*2)
	}
}

func main() {
	//声明一个10长度的数组
	//var a[10]int
	//a[0] = 100
	//a[1] = 1
	//arr2 := a
	//arr2[0] = 222
	//fmt.Println("aaaa",a)
	//fmt.Println("aaaa",arr2)
	test()
	//str := "Go is a beautiful language!"
	//fmt.Printf("The length of str is: %d\n", len(str))
	//for pos, char := range str {
	//	fmt.Printf("Character on position %d is: %c \n", pos, char)
	//}
}
