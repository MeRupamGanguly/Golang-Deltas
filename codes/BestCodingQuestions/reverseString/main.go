package main

import "fmt"

func reverseString(s string) string {
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	for l < r {
		t := arr[l]
		arr[l] = arr[r]
		arr[r] = t
		l++
		r--
	}
	return string(arr)
}
func main() {
	fmt.Println(reverseString("Hello"))
}
