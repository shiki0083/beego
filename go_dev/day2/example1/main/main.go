package main

import "fmt"

func list(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i, n-i, n)
	}
}

func main() {
	list(5)
}
