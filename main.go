package main

import "fmt"

func reverseString(s string) string {
	var result string

	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func main() {
	fmt.Println(reverseString("hello"))
}
