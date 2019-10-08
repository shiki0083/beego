package main

import (
	"fmt"
	"strings"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = "http://" + url
	}
	return url
}

func main() {
	var (
		url  string = "192.168.0.1.1"
		path string
	)

	url = urlProcess(url)
	//pathProcess(path)
	fmt.Println("%ss", url)
	fmt.Println("%ss", path)
}
