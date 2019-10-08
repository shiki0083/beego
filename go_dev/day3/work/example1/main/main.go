package main

import "fmt"

// 获取素数
func isOk(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 100
	m := 200
	for i := n; i < m; i++ {
		if isOk(i) == true {
			fmt.Println("hh", i)
			continue
		}
	}
}
