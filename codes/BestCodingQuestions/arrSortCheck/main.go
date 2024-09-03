package main

import "fmt"

func arraySortCheck(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(arraySortCheck([]int{2, 4, 6, 8, 1}))
}
