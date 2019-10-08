package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute"
)

//利用go 的goroute 实现不同文件的并发操作
//通过channel 管道（chan） 线程数据共同
func main() {
	//声明一个管道pipe
	var pipe chan int
	pipe = make(chan int, 3)
	go goroute.Add(100, 200, pipe)
	//把管道里的值提到sum上
	sum := <-pipe
	fmt.Println("sum==", sum)

}
