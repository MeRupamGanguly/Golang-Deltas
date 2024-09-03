package main

import "fmt"

func findIndex(a []int, inp int) int {
	for i := range a {
		if a[i] == inp {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(findIndex([]int{2, 4, 6, 8, 12, 24, 25, 27, 34, 45, 67, 32, 47, 89}, 8))
}
