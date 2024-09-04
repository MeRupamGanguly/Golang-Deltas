package main

import "fmt"

func palindrome(s string) bool {
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
func main() {
	fmt.Println(palindrome("hello"))
}
