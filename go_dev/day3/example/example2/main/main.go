package main

import "fmt"

func main()  {
	var a int = 10
	fmt.Println("ssss",&a)
	// 指针变量声明 指针存的是地址！ 指针就是一个存放地址的变量
	var p *int
	p = &a
	fmt.Println("aaaa",p)
	// *指针就能取出对应的值
	*p = 100
	fmt.Println("ssss",a)
}
