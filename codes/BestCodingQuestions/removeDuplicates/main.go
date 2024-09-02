package main

import "fmt"

func removeDup(arr []int) []int {
	l := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[l] {
			l++
			arr[l] = arr[i]
		}
	}
	return arr[:l+1]
}
func removeDuplicates(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[i+1] {
			temp := arr[i+1:]
			arr = arr[:i]
			arr = append(arr, temp...)
		}
	}
	return arr
}
func main() {
	arr := []int{2, 4, 5, 6, 6, 8, 8, 9, 10, 11, 12, 12, 13, 13}
	fmt.Println(removeDup(arr))
}
