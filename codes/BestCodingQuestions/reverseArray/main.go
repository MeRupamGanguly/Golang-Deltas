package main

import "fmt"

func reverseArray(arr []int) []int {
	l := 0
	r := len(arr) - 1
	for l < r {
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		l++
		r--
	}
	return arr
}
func main() {
	arr := []int{2, 3, 8, 6, 9, 11, 15, 13, 12}
	fmt.Println(reverseArray(arr))
}
