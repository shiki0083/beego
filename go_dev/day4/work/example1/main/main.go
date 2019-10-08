package main

func process(str string)bool  {
	for i := 0;i<len(str) ;i++  {
		if i == len(str)/2 {
			break
		}
		last := len(str) -i -1
		if str[i] != str[last] {
			return false
		}
	}
	return true
}

func main()  {
	
}